package routes

import (
	"blog/app/web/controllers"
	adminv1 "blog/app/web/controllers/admin/v1"
	"blog/app/web/middlewares"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func InitRouter(application *iris.Application) {
	// 注册后台接口路由
	adminRoute(application)
	// 注册 api 接口
	ApiRoute(application)
	// 注册 web 路由
	webRoute(application)
}

func adminRoute(api *iris.Application) {
	api.Get("/api/admin", func(context iris.Context) {
		context.WriteString("hello admin api")
	})
	// 登录相关接口
	login := api.Party("/api/admin/v1")
	{
		login.Post("/login", adminv1.Login)
		login.Post("/logout", adminv1.Logout).Use(middlewares.JwtHandler().Serve)
	}
	mvc.Configure(api.Party("/api/admin/v1"), func(v1 *mvc.Application) {
		v1.Router.Use(middlewares.JwtHandler().Serve)
		// 用户相关
		v1.Party("/user").Handle(adminv1.NewUserController())
	})
}

func ApiRoute(application *iris.Application) {
	mvc.New(application.Party("/api/test")).Handle(controllers.NewTestController())
}

func webRoute(application *iris.Application) {
	// hello
	application.Get("/", func(ctx iris.Context) {
		ctx.WriteString("hello blog")
	})
}
