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

var data = map[string]interface{}{
	"name":        "test",
	"description": "测试",
}

func TestCategoryList(t *testing.T) {
	app := bootstrap.Register()
	test := httptest.New(t, app)
	token, err := jwt.MakeToken(1, "admin", "1986410806@qq.com", "管理员")
	if err != nil {
		t.Error(err)
	}

	page := map[string]int{
		"page":  1,
		"limit": 10,
	}
	// 列表
	test.GET("/admin/v1/category/list").
		WithHeader("Authorization", "Bearer "+token).
		WithQueryObject(page).
		Expect().Status(httptest.StatusOK).
		JSON().Object().
		ValueEqual("errorCode", 0).
		ValueEqual("success", true)

}

func TestCategoryCreate(t *testing.T) {
	app := bootstrap.Register()
	test := httptest.New(t, app)
	token, err := jwt.MakeToken(1, "admin", "1986410806@qq.com", "管理员")
	if err != nil {
		t.Error(err)
	}
	// 新增
	test.POST("/admin/v1/category/create").
		WithHeader("Authorization", "Bearer "+token).WithForm(data).
		Expect().Status(httptest.StatusOK).
		JSON().Object().
		ValueEqual("errorCode", 0).
		ValueEqual("success", true)
}

func TestCategoryUpdate(t *testing.T) {
	app := bootstrap.Register()
	test := httptest.New(t, app)
	token, err := jwt.MakeToken(1, "admin", "1986410806@qq.com", "管理员")
	if err != nil {
		t.Error(err)
	}
	out := repositories.NewCategoryRepository(database.DB()).GetByName(data["name"].(string))

	data["description"] = "测试更新"
	test.POST("/admin/v1/category/update").
		WithHeader("Authorization", "Bearer "+token).
		WithForm(data).
		WithFormField("id", out.ID).
		Expect().Status(httptest.StatusOK).
		JSON().Object().
		ValueEqual("errorCode", 0).
		ValueEqual("success", true)
}

func TestCategoryDel(t *testing.T) {
	app := bootstrap.Register()
	test := httptest.New(t, app)
	token, err := jwt.MakeToken(1, "admin", "1986410806@qq.com", "管理员")
	if err != nil {
		t.Error(err)
	}
	out := repositories.NewCategoryRepository(database.DB()).GetByName(data["name"].(string))
	// 删除
	test.POST("/admin/v1/category/del").
		WithHeader("Authorization", "Bearer "+token).
		WithFormField("id", out.ID).
		Expect().Status(httptest.StatusOK).
		JSON().Object().
		ValueEqual("errorCode", 0).
		ValueEqual("success", true)
}
