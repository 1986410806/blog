package repositories

import (
	"blog/app/models"
	"blog/database"
	"errors"
	"github.com/jinzhu/gorm"
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
	var tag = this.GetById(id)
	if tag == nil {
		return nil, errors.New("标签不存在")
	}
	err := this.db.Model(&tag).Updates(data).Error
	return tag, err
}

// 使用 model 编辑标签
func (this *TagRepositories) UpdateByModel(tag *models.Tag, data UpdateData) (err error) {
	err = this.db.Model(&tag).Updates(data).Error
	return err
}

// 使用id 删除标签
func (this *TagRepositories) DelById(id int) (err error) {
	var tag models.Tag
	return database.DB().Delete(tag,id).Error
}

// 按id查找
func (this *TagRepositories) GetById(id int) (*models.Tag) {
	var tag = &models.Tag{}
	this.db.First(tag, uint(id))
	return tag
}

// 按 name 查找
func (this *TagRepositories) GetByName(name string) (*models.Tag) {
	var tag = &models.Tag{}
	this.db.Where("name = ?", name).First(tag)
	return tag
}
