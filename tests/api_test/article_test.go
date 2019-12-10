package api

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
func TestArticleCreate(t *testing.T) {
	app := bootstrap.Register()
	test := httptest.New(t, app)
	token, err := jwt.MakeToken(1, "admin", "1986410806@qq.com", "管理员")
	if err != nil {
		t.Error(err)
	}
	data := map[string]interface{}{
		"tags":        "测试,ceshi",
		"title":       "测试发布文章",
		"summary":     "这是一个测试文章",
		"content":     "这是一个测试文章这是一个测试文章这是一个测试文章这是一个测试文章",
		"category_id": 1,
	}
	// 新增
	test.POST("/admin/v1/article/create").
		WithHeader("Authorization", "Bearer "+token).WithForm(data).
		Expect().Status(httptest.StatusOK).
		JSON().Object().
		ValueEqual("errorCode", 0).
		ValueEqual("success", true)
}

// 编辑文章
func TestArticleUpdate(t *testing.T) {
	app := bootstrap.Register()
	test := httptest.New(t, app)
	token, err := jwt.MakeToken(1, "admin", "1986410806@qq.com", "管理员")
	if err != nil {
		t.Error(err)
	}
	data := map[string]interface{}{
		"article_id":  1,
		"tags":        "测试,ceshi,hahah",
		"title":       "测试发布文章1",
		"summary":     "这是一个测试文章2",
		"content":     "这是一个测试文章这是一个测试文章这是一个测试文章这是一个测试文章",
		"category_id": 1,
	}
	// 新增
	test.POST("/admin/v1/article/update").
		WithHeader("Authorization", "Bearer "+token).WithForm(data).
		Expect().Status(httptest.StatusOK).
		JSON().Object().
		ValueEqual("errorCode", 0).
		ValueEqual("success", true)
}

// 获取文章列表
func TestArticleList(t *testing.T) {
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
	test.GET("/admin/v1/article/list").
		WithHeader("Authorization", "Bearer "+token).
		WithQueryObject(page).
		Expect().Status(httptest.StatusOK).
		JSON().Object().
		ValueEqual("errorCode", 0).
		ValueEqual("success", true)
}
