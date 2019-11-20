package services

import (
	"blog/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func NewUserTokenService() *UserTokenService {
	return &UserTokenService{}
}

type UserTokenService struct {
}

func (userToken UserTokenService) MakeToken(uid uint) (string, error) {
	//生成加密串过程
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"nick_name": "iris",
			"email":     "go-iris@qq.com",
			"id":        uid,
			"iss":       "Iris",
			"iat":       time.Now().Unix(),
			"jti":       "9527",
			"exp":       time.Now().Add(10 * time.Hour * time.Duration(1)).Unix(),
		})

	return token.SignedString([]byte(config.Conf.JwtSecret))
}
