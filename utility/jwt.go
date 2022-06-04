package utility

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//validity 有效
const validity = time.Hour

type jwtToken struct {
	Token        string
	Exp          time.Time
	RefreshToken string
}

func NewJwtToken(secret []byte, uid int) (*jwtToken, error) {
	t := jwt.New(jwt.SigningMethodHS256)

	exp := time.Now().Add(validity) // 设置超时时间

	claims := make(jwt.MapClaims)
	claims["exp"] = exp.Unix()
	claims["iat"] = time.Now().Unix()
	claims["uid"] = uid

	t.Claims = claims

	token, err := t.SignedString(secret)

	if err != nil {
		return nil, fmt.Errorf("gen token err: %s", err.Error())
	}

	return &jwtToken{
		Token: token,
		Exp:   exp,
	}, nil
}

func ParseToken(tokenString string, secret []byte) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("token is unValid")
	}
}
