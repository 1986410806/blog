package bootstrap

import (
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

func Run() {
	app := iris.New()
	// 初始化数据库
	database.OpenDB(config.Conf.MySqlUrl)

	// 初始化日志
	initLogrus()
	// 初始化路由
	routes.InitRouter(app)

	app.AllowMethods(iris.MethodOptions)
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

	if err != nil {
		panic(err)
	}
}

// 中间件注册
func registerMiddleware(app *iris.Application) {

	app.Configure(iris.WithOptimizations)
	// 注册 header 跨域设置
	app.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	}))
	//并将请求记录到终端。
	app.Use(recover.New())

	app.Use(logger.New())
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
