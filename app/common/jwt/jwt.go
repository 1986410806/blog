package jwt

import (
	"blog/config"
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple"
	"time"
)

var tokenClaimsData *tokenClaims

type tokenClaims struct {
	UserId   uint
	UserName string
}

func JwtHandler() *jwtmiddleware.Middleware {
	return jwtmiddleware.New(jwtmiddleware.Config{
		// 这个方法将验证jwt的token
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Conf.JwtSecret), nil
		},
		//设置后，中间件会验证令牌是否使用特定的签名算法进行签名
		//如果签名方法不是常量，则可以使用ValidationKeyGetter回调来实现其他检查
		//重要的是要避免此处的安全问题：https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		//加密的方式
		SigningMethod: jwt.SigningMethodHS256,
		//验证未通过错误处理方式
		ErrorHandler: func(ctx iris.Context, err error) {
			println(err.Error())
			ctx.JSON(simple.JsonErrorMsg(err.Error()))
		},
		Expiration: true,
	})

}

func GetTokenClaim(Ctx iris.Context) *tokenClaims {
	if tokenClaimsData == nil {
		claims := Ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
		tokenClaimsData = &tokenClaims{
			UserId:   uint(claims["id"].(float64)),
			UserName: claims["user_name"].(string),
		}
	}
	return tokenClaimsData
}

func MakeToken(uid uint, userName, email string) (string, error) {
	//生成加密串过程
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_name": userName,
			"email":     email,
			"id":        uid,
			"iss":       "blog-admin",
			"iat":       time.Now().Unix(),
			"jti":       simple.Uuid(),
			"exp":       time.Now().Add(10 * time.Hour * time.Duration(1)).Unix(),
		})

	return token.SignedString([]byte(config.Conf.JwtSecret))
}
