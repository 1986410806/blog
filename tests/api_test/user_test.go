package api_test

import (
	"blog/app/common/jwt"
	"blog/bootstrap"
	"blog/config"
	"github.com/kataras/iris/v12/httptest"
	"testing"
)

func init() {
	config.InitConfig("../../blog.yaml")
}

func TestUserInfo(t *testing.T) {
	app := bootstrap.Register()
	e := httptest.New(t, app)
	token, err := jwt.MakeToken(1, "admin", "1986410806@qq.com")
	if err != nil {
		t.Error(err)
	}

	e.GET("/api/admin/v1/user").
		WithHeader("Authorization", "Bearer "+token).
		Expect().Status(httptest.StatusOK).
		JSON().Object().
		ValueEqual("errorCode", 0).
		ValueEqual("success", true)
}
