package dao

import (
	"context"
	"fmt"
	"testing"
)

func TestUserDao_GetUserByEmail(t *testing.T) {
	u, _ := User.GetUserByEmail(context.Background(), "1986410806@qq.com")

	fmt.Println(u)
}
