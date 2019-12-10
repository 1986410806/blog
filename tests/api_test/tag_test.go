package api

import (
	"blog/app/common/jwt"
	"blog/app/repositories"
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
	token, err := jwt.MakeToken(1, "admin", "1986410806@qq.com", "管理员")
	if err != nil {
		t.Error(err)
	}
	data := map[string]interface{}{
		"name":        "test",
		"description": "测试",
	}
	// 新增
	test.POST("/admin/v1/tag/create").
		WithHeader("Authorization", "Bearer "+token).WithForm(data).
		Expect().Status(httptest.StatusOK).
		JSON().Object().
		ValueEqual("errorCode", 0).
		ValueEqual("success", true)
	out := repositories.NewTagRepository(database.DB()).GetByName(data["name"].(string))
	// 列表
	page := map[string]int{
		"page":  1,
		"limit": 10,
	}
	test.GET("/admin/v1/tag/list").
		WithHeader("Authorization", "Bearer "+token).
		WithQueryObject(page).
		Expect().Status(httptest.StatusOK).
		JSON().Object().
		ValueEqual("errorCode", 0).
		ValueEqual("success", true)

	// 更新
	data["description"] = "测试更新"
	test.POST("/admin/v1/tag/update").
		WithHeader("Authorization", "Bearer "+token).
		WithForm(data).
		WithFormField("id", out.ID).
		Expect().Status(httptest.StatusOK).
		JSON().Object().
		ValueEqual("errorCode", 0).
		ValueEqual("success", true)
	// 删除
	test.POST("/admin/v1/tag/del").
		WithHeader("Authorization", "Bearer "+token).
		WithFormField("id", out.ID).
		Expect().Status(httptest.StatusOK).
		JSON().Object().
		ValueEqual("errorCode", 0).
		ValueEqual("success", true)
}
