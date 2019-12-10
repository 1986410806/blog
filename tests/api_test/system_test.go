package api

import (
	"blog/app/common/jwt"
	"blog/app/repositories"
	"blog/bootstrap"
	"blog/config"
	"github.com/kataras/iris/v12/httptest"
	"testing"
)

func init() {
	config.InitConfig("../../blog.yaml")
}

// 增删改查
func TestSystemConfigCurd(t *testing.T) {
	app := bootstrap.Register()
	test := httptest.New(t, app)
	token, err := jwt.MakeToken(1, "admin", "1986410806@qq.com", "管理员")
	if err != nil {
		t.Error(err)
	}
	data := map[string]interface{}{
		"key":         "test",
		"value":       "test",
		"name":        "测试key",
		"description": "测试",
	}
	// 新增
	test.POST("/admin/v1/system/config/create").
		WithHeader("Authorization", "Bearer "+token).WithForm(data).
		Expect().Status(httptest.StatusOK).
		JSON().Object().
		ValueEqual("errorCode", 0).
		ValueEqual("success", true)
	// 列表
	out := repositories.NewSystemConfigRepository().GetByKey(data["key"].(string))

	page := map[string]int{
		"page":  1,
		"limit": 10,
	}
	test.GET("/admin/v1/system/config/list").
		WithHeader("Authorization", "Bearer "+token).
		WithQueryObject(page).
		Expect().Status(httptest.StatusOK).
		JSON().Object().
		ValueEqual("errorCode", 0).
		ValueEqual("success", true)
	// 更新
	data["description"] = "测试更新"

	test.POST("/admin/v1/system/config/update").
		WithHeader("Authorization", "Bearer "+token).
		WithForm(data).
		WithFormField("id", out.ID).
		Expect().Status(httptest.StatusOK).
		JSON().Object().
		ValueEqual("errorCode", 0).
		ValueEqual("success", true)
	// 删除
	test.POST("/admin/v1/system/config/del").
		WithHeader("Authorization", "Bearer "+token).
		WithFormField("id", out.ID).
		Expect().Status(httptest.StatusOK).
		JSON().Object().
		ValueEqual("errorCode", 0).
		ValueEqual("success", true)
}

// 系统配置
func TestConfig(t *testing.T) {
	app := bootstrap.Register()
	test := httptest.New(t, app)
	test.GET("/admin/v1/system/config").
		Expect().Status(httptest.StatusOK).
		JSON().Object().
		ValueEqual("errorCode", 0).
		ValueEqual("success", true)

}
