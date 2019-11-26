package models

import (
	"github.com/jinzhu/gorm"
)

// 标签
type Tag struct {
	gorm.Model
	Name        string `gorm:"size:32;unique;not null" json:"name" form:"name"`
	Description string `gorm:"size:1024" json:"description" form:"description"`
	Status      int    `gorm:"not null" json:"status" form:"status"`
}