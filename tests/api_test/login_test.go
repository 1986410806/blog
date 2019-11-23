package api_test

/**
 * 测试源码
 * @like https://github.com/iris-contrib/httpexpect
 * 文档
 * @like https://iris-go.com/start/#testing
 */
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

func TestLogin(t *testing.T) {
	app := bootstrap.Register()
	e := httptest.New(t, app)
	// 登录
	loginData := map[string]string{
		"username": "admin",
		"password": "123456",
		"ref":      "aaa",
	}
	e.POST("/api/admin/v1/login").
		WithForm(loginData).Expect().
		Status(httptest.StatusOK).
		JSON().
		Object().
		ValueEqual("errorCode", 0).
		ValueEqual("success", true)

}

func TestLogout(t *testing.T) {
	app := bootstrap.Register()
	e := httptest.New(t, app)
	token, err := jwt.MakeToken(1, "admin", "1986410806@qq.com")
	if err != nil {
		t.Error(err)
	}

	e.POST("/api/admin/v1/logout").
		WithHeader("Authorization", "Bearer "+token).
		Expect().
		Status(httptest.StatusOK).
		JSON().
		Object().
		ValueEqual("errorCode", 0).
		ValueEqual("success", true)
}
