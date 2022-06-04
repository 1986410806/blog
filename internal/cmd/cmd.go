package cmd

import (
	"blog-go/internal/middleware"
	"context"

	"blog-go/internal/controller"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			s.BindMiddlewareDefault(ghttp.MiddlewareHandlerResponse, middleware.CORS)
			s.BindObject("/", controller.Auth)
			s.BindObject("/", controller.Hello)

			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(middleware.JwtHandler)
				group.Bind(
					controller.User,
					controller.Tag,
					controller.File,
				)

			})
			s.Run()
			return nil
		},
	}
)
