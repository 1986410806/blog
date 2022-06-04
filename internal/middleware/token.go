package middleware

import (
	"blog-go/utility"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
	"net/http"
	"strings"
)

//goland:noinspection GoUnusedExportedFunction
func JwtHandler(r *ghttp.Request) {

	token, err := getToken(r)
	if err != nil {
		glog.Error(r.Context(), err)
		r.Response.WriteStatusExit(http.StatusForbidden)
	}
	secret, err := g.Cfg().Get(r.Context(), "token.secret")
	if err != nil {
		glog.Error(r.Context(), err)
		r.Response.WriteStatusExit(http.StatusInternalServerError)
	}
	jwtClaims, err := utility.ParseToken(token, secret.Bytes())
	if err != nil {
		r.Response.WriteStatusExit(http.StatusForbidden)
	}

	if err != nil {
		glog.Error(r.Context(), err)
		r.Response.WriteStatusExit(http.StatusForbidden)
	}

	if jwtClaims == nil {
		r.Response.WriteStatusExit(http.StatusForbidden)
		return
	}
	r.SetParam("uid", jwtClaims["uid"])
	r.Middleware.Next()
}

func getToken(r *ghttp.Request) (string, error) {

	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		return "", errors.New("token not empty")
	}

	parts := strings.SplitN(authHeader, " ", 2)

	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", errors.New("auth header is invalid")
	}

	return parts[1], nil
}
