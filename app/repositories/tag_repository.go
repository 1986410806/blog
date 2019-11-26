package repositories

import (
	"blog/app/models"
	"blog/database"
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type TagRepositories struct {
	db *gorm.DB
}

func NewTagRepositories() *TagRepositories {
	return &TagRepositories{
		db: database.DB(),
	}
}

// 获取标签列表
func (this *TagRepositories) TagList() (list []models.Tag, err error) {
	err = this.db.Find(&list).Error
	return list, err
}

// 创建标签
func (this *TagRepositories) Create(tag *models.Tag) (*models.Tag, error) {

	if tmp := this.GetByName(tag.Name); tmp.Name == tag.Name {
		return tmp, nil
	}
	err := this.db.Create(&tag).Error
	return tag, err
}

// 编辑标签
func (this *TagRepositories) UpdateById(id int, data UpdateData) (*models.Tag, error) {
	var tag, err = this.GetById(id)
	if err != nil {
		return nil, errors.New("标签不存在")
	}
	err = this.db.Model(&tag).Updates(data).Error
	return tag, err
}

// 使用 model 编辑标签
func (this *TagRepositories) UpdateByModel(tag *models.Tag, data UpdateData) (err error) {
	err = this.db.Model(&tag).Updates(data).Error
	return err
}

// 使用id 删除标签
func (this *TagRepositories) DelById(id int) error {
	var tag, err = this.GetById(id)
	if err != nil {
		return errors.New("标签不存在")
	}
	// 为避免唯一索引冲突 软删除的时候给唯一列拼接上 当前时间
	// 也可以数据库上使用和软删除字端组合唯一列
	return this.db.Model(&tag).Updates(UpdateData{
		"deleted_at": time.Now(),
		"name":       tag.Name + time.Now().Format("2006-01-02 15:04:05"),
	}).Error
}

// 按id查找
func (this *TagRepositories) GetById(id int) (*models.Tag, error) {
	var tag = &models.Tag{}
	e := this.db.First(tag, uint(id)).Error
	return tag, e
}

// 按 name 查找
func (this *TagRepositories) GetByName(name string) (*models.Tag) {
	var tag = &models.Tag{}
	this.db.Where("name = ?", name).First(tag)
	return tag
}
