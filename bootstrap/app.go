package bootstrap

import (
	"blog/config"
	"blog/database"
	"blog/routes"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/sirupsen/logrus"
	"os"
)

func Run() {
	app := iris.New()
	// 初始化数据库
	err := database.OpenDB(config.Conf.MySqlUrl)
	if err != nil {
		panic(err)
	}
	// 初始化日志
	initLogrus()
	// 初始化路由
	routes.InitRouter(app)

	app.AllowMethods(iris.MethodOptions)

	err = app.Run(iris.Addr(":"+string(config.Conf.Port)), iris.WithoutServerError(iris.ErrServerClosed))
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
