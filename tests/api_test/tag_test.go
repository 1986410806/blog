package api

import (
	"blog/app/common/jwt"
	"blog/bootstrap"
	"blog/config"
	"blog/database"
	"github.com/kataras/iris/v12/httptest"
	"testing"
)

func init() {
	config.InitConfig("../../blog.yaml")
}

// 增删改查
func TestTagCurd(t *testing.T) {
	app := bootstrap.Register()
	test := httptest.New(t, app)
	token, err := jwt.MakeToken(1, "admin", "1986410806@qq.com")
	if err != nil {
		t.Error(err)
	}
	data := map[string]interface{}{
		"name":        "test",
		"description": "测试",
	}
	database.DB().Table("t_tag").Where("name=?", data["name"]).UpdateColumn("deleted_at", nil)
	// 新增
	test.POST("/api/admin/v1/tag").
		WithHeader("Authorization", "Bearer "+token).WithForm(data).
		Expect().Status(httptest.StatusOK).
		JSON().Object().
		ValueEqual("errorCode", 0).
		ValueEqual("success", true)
	// 列表
	test.GET("/api/admin/v1/tag/list").
		WithHeader("Authorization", "Bearer "+token).
		Expect().Status(httptest.StatusOK).
		JSON().Object().
		ValueEqual("errorCode", 0).
		ValueEqual("success", true)
	// 更新
	test.POST("/api/admin/v1/tag/update").
		WithHeader("Authorization", "Bearer "+token).
		WithForm(data).
		WithFormField("id", 1).
		Expect().Status(httptest.StatusOK).
		JSON().Object().
		ValueEqual("errorCode", 0).
		ValueEqual("success", true)
	// 删除
	test.POST("/api/admin/v1/tag/del").
		WithHeader("Authorization", "Bearer "+token).
		WithFormField("id", 1).
		Expect().Status(httptest.StatusOK).
		JSON().Object().
		ValueEqual("errorCode", 0).
		ValueEqual("success", true)
}
