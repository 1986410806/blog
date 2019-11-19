package routes

import (
	"blog/app/web/controllers"
	adminv1 "blog/app/web/controllers/admin/v1"
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

func adminRoute(application *iris.Application) {
	application.Get("/api/admin", func(context iris.Context) {
		context.WriteString("hello admin api")
	})
	mvc.Configure(application.Party("/api/admin/v1"), func(v1 *mvc.Application) {
		// 登录相关接口
		v1.Handle(adminv1.NewLoginController())
	})
}

func ApiRoute(application *iris.Application) {
	mvc.New(application.Party("/api/v1/test")).Handle(controllers.NewTestController())
}

func webRoute(application *iris.Application) {
	// hello
	application.Get("/", func(ctx iris.Context) {
		ctx.WriteString("hello blog")
	})
}
