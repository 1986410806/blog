package routes

import (
	controllers "blog/app/controllers"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func InitRouter(application *iris.Application) {
	basePath := "api/v1"
	// test
	application.Get("/", func(ctx iris.Context) {
		ctx.WriteString("hello web")
	})
	// mvc test
	mvc.New(application.Party(basePath + "/test")).Handle(controllers.NewTestController())
}

func adminRoute(application *iris.Application) {
	mvc.New(application.Party(basePath + "/test")).Handle(controllers.NewTestController())

}

func ApiRoute(application *iris.Application) {

}
