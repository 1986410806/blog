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
	api.Get("/admin", func(context iris.Context) {
		context.WriteString("hello admin api")
	})

	v1 := api.Party("/admin/v1")
	{
		// 登录相关接口
		v1.Post("/login", adminv1.Login)
		v1.Post("/logout", adminv1.Logout).Use(middlewares.JwtHandler)
		mvc.New(v1.Party("/system")).Handle(adminv1.NewSystemController())
		// 需要登录授权的
		mvc.Configure(v1.Party("/", middlewares.JwtHandler), func(auth *mvc.Application) {
			// 应用配置
			auth.Party("/system/config").Handle(adminv1.NewSystemConfigController())
			// 用户相关
			auth.Party("/user").Handle(adminv1.NewUserController())
			// 栏目
			auth.Party("/category").Handle(adminv1.NewCategoryController())
			// 标签
			auth.Party("/tag").Handle(adminv1.NewTagController())
			// 文章
			auth.Party("/article").Handle(adminv1.NewArticleController())
		})
	}

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
