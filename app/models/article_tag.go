package models

import "github.com/jinzhu/gorm"

const (
	ArticleTagStatusOk      = 0
	ArticleTagStatusDeleted = 1
)

// 文章标签
type ArticleTag struct {
	gorm.Model
	ArticleId int64 `gorm:"not null;index:idx_article_id;" json:"articleId" form:"articleId"` // 文章编号
	TagId     int64 `gorm:"not null;index:idx_tag_id;" json:"tagId" form:"tagId"`             // 标签编号
	Status    uint8 `gorm:"not null;index:idx_status;default:0" json:"status" form:"status"`  // 状态：正常、删除
	// 关联
	Tag *Tag `gorm:"foreignkey:TagId"`
}
