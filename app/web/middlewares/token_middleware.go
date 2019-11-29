package middlewares

import (
	"blog/app/common/jwt"
	"github.com/kataras/iris/v12/context"
)

// jwt 验证注册
func JwtHandler(ctx context.Context) {
	jwt.JwtHandler(ctx)
}
