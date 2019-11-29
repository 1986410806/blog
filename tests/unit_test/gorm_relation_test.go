package unit

import (
	"blog/app/models"
	"blog/bootstrap"
	"blog/config"
	"blog/database"
	"testing"
)

func init() {
	config.InitConfig("../../blog.yaml")
	bootstrap.Register()
}

// 模型 一对一 测试
func TestHasOne(t *testing.T) {
	db := database.DB()
	var article = &models.Article{}
	db = db.Preload("Category").
		Preload("User").
		First(article)
	switch {
	case db.Error != nil:
		t.Error(db.Error)
	case db.RecordNotFound():
		t.Error("记录为空")
	}
}

// 模型 一对多 测试
func TestHasMany(t *testing.T) {
	db := database.DB()
	var article = &models.Article{}
	db = db.Preload("ArticleTag").
		Preload("ArticleTag.Tag").
		First(article)
	switch {
	case db.Error != nil:
		t.Error(db.Error)
	case db.RecordNotFound():
		t.Error("记录为空")
	}
}
