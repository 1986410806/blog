package common

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// 加密
func EncodePassword(rawPassword string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	return string(hash)
}

// 验证密码
func ValidatePassword(encodePassword, inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encodePassword),
		[]byte(inputPassword))
	return err == nil
}
