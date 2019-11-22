package api_test

import (
	"blog/bootstrap"
	"blog/config"
	"github.com/kataras/iris/v12/httptest"
	"testing"
)

func init() {
	config.InitConfig("../blog.yaml")
}

func TestUserInfo(t *testing.T) {
	//app := bootstrap.Register()
	//e := httptest.New(t, app)
	//token :=
	//e.POST("/api/admin/v1/user").
	//	WithHeader("Authorization",)
	//	Expect().
	//	Status(httptest.StatusOK).
	//	JSON().
	//	Object().
	//	ValueEqual("errorCode", 0).
	//	ValueEqual("success", true)
}
