package middlewares

import (
	"blog/config"
	"github.com/betacraft/yaag/irisyaag"
	"github.com/betacraft/yaag/yaag"
	"github.com/kataras/iris/v12/context"
)

// jwt 验证注册
func ApiDocHandler() context.Handler {
	yaag.Init(&yaag.Config{ // <- IMPORTANT, init the middleware.
		On:       config.Conf.Env == "local",
		DocTitle: "Iris",
		DocPath:  config.Conf.ApiDoc.Path + "admin.html",
		BaseUrls: map[string]string{
			"Production": config.Conf.ApiDoc.ProdUrl,
			"Staging":    config.Conf.ApiDoc.DevUrl},
	})
	return irisyaag.New() // <- IMPORTANT, register the middleware.
}
