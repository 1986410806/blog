package models

import "github.com/jinzhu/gorm"

const (
	ArticleStatusPublished = 0 // 已发布
	ArticleStatusDeleted   = 1 // 已删除
	ArticleStatusDraft     = 2 // 草稿

	TopicStatusOk      = 0
	TopicStatusDeleted = 1

	TopicTagStatusOk      = 0
	TopicTagStatusDeleted = 1

	ContentTypeHtml     = "html"
	ContentTypeMarkdown = "markdown"
)

// 文章
type Article struct {
	gorm.Model
	CategoryId  int64  `gorm:"index:idx_category_id;not null" json:"categoryId" form:"categoryId"` // 分类编号
	UserId      uint   `gorm:"index:idx_user_id;default:0" json:"userId" form:"userId"`            // 所属用户编号
	Title       string `gorm:"size:128;not null;" json:"title" form:"title"`                       // 标题
	Summary     string `gorm:"type:text" json:"summary" form:"summary"`                            // 摘要
	Content     string `gorm:"type:longtext;not null;" json:"content" form:"content"`              // 内容
	ContentType string `gorm:"type:varchar(32);not null" json:"contentType" form:"contentType"`    // 内容类型：markdown、html
	Status      int    `gorm:"int;not null" json:"status" form:"status"`                           // 状态
	Share       bool   `gorm:"not null" json:"share" form:"share"`                                 // 是否是分享的文章，如果是这里只会显示文章摘要，原文需要跳往原链接查看
	SourceUrl   string `gorm:"type:text" json:"sourceUrl" form:"sourceUrl"`                        // 原文链接
	// 关联
	User       User     `gorm:"foreignkey:UserId"`
	Category   Category `gorm:"foreignkey:CategoryId"`
	ArticleTag []ArticleTag
}
