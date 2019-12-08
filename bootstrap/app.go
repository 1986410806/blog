package bootstrap

import (
	"blog/app/web/middlewares"
	"blog/config"
	"blog/database"
	"blog/routes"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/mlogclub/simple"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// 注册 web 程序所需要 服务
func Register() *iris.Application {
	app := iris.New()
	// 初始化数据库
	database.OpenDB(config.Conf.MySqlUrl)
	// redis 初始化
	database.RedisInit()
	// 注册中间件
	registerMiddleware(app)
	// 初始化日志
	initLogrus()
	// 初始化路由
	routes.InitRouter(app)

	return app
}

// 运行 web 程序
func Run(app *iris.Application) error {

	server := &http.Server{Addr: ":" + config.Conf.Port}

	handleSignal(server)

	err := app.Run(iris.Server(server), iris.WithConfiguration(iris.Configuration{
		DisableStartupLog:                 false,
		DisableInterruptHandler:           false,
		DisablePathCorrection:             false,
		EnablePathEscape:                  false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: false,
		DisableAutoFireStatusCode:         false,
		EnableOptimizations:               true,
		TimeFormat:                        "2006-01-02 15:04:05",
		Charset:                           "UTF-8",
	}))
	return err
}

// 中间件注册
func registerMiddleware(app *iris.Application) {

	app.AllowMethods(iris.MethodOptions)

	app.Configure(iris.WithOptimizations)
	// 注册 header 跨域设置
	app.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	}))
	// apidoc
	app.Use(middlewares.ApiDocHandler())

	if config.Conf.Debug == true {
		app.Use(recover.New())
		//
		app.Use(logger.New(logger.Config{
			//状态显示状态代码
			Status: true,
			// IP显示请求的远程地址
			IP: true,
			//方法显示http方法
			Method: true,
			// Path显示请求路径
			Path: true,
			// Query将url查询附加到Path。
			Query: true,
			//Columns：true，
			// 如果不为空然后它的内容来自`ctx.Values(),Get("logger_message")
			//将添加到日志中。
			MessageContextKeys: []string{"logger_message"},
			//如果不为空然后它的内容来自`ctx.GetHeader（“User-Agent”）
			MessageHeaderKeys: []string{"User-Agent"},
		}))
	}
}

// 初始化日志
func initLogrus() {
	file, err := os.OpenFile(config.Conf.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logrus.SetOutput(file)
	} else {
		logrus.Error(err)
	}
}

func handleSignal(server *http.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		s := <-c
		logrus.Infof("got signal [%s], exiting now", s)
		if err := server.Close(); nil != err {
			logrus.Errorf("server close failed: " + err.Error())
		}

		simple.CloseDB()

		logrus.Infof("Exited")
		os.Exit(0)
	}()
}
