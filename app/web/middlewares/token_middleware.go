package middlewares

import (
	"blog/app/common/jwt"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
)

func JwtHandler() *jwtmiddleware.Middleware {
	return jwt.JwtHandler()
}
