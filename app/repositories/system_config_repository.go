package repositories

import (
	"blog/app/models"
	"blog/database"
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type SystemConfigRepositories struct {
	db *gorm.DB
}

func NewSystemConfigRepositories() *SystemConfigRepositories {
	return &SystemConfigRepositories{
		db: database.DB(),
	}
}

// 获取配置列表
func (this *SystemConfigRepositories) ConfigList() (list []models.SysConfig, err error) {
	err = this.db.Find(&list).Error
	return list, err
}

// 创建配置
func (this *SystemConfigRepositories) Create(model *models.SysConfig) (*models.SysConfig, error) {
	if tmp := this.GetByKey(model.Key); tmp.Key == model.Key {
		return tmp, nil
	}
	err := this.db.Create(&model).Error
	return model, err
}

// 编辑配置
func (this *SystemConfigRepositories) UpdateById(id int, data UpdateData) (*models.SysConfig, error) {
	model, err := this.GetById(id)
	if err != nil {
		return nil, errors.New("配置不存在")
	}
	db := this.db.Model(&model).Updates(data)
	// 判断影响的行数
	if db.RowsAffected < 1 {
		return nil, errors.New("更新数据失败～～")
	}
	return model, db.Error
}

// 使用 model 编辑配置
func (this *SystemConfigRepositories) UpdateByModel(model *models.SysConfig, data UpdateData) error {
	db := this.db.Model(&model).Updates(data)
	// 判断影响的行数
	if db.RowsAffected > 0 {
		return errors.New("更新数据失败～～")
	}
	return db.Error
}

// 使用id 删除配置
func (this *SystemConfigRepositories) DelById(id int) (err error) {
	model, err := this.GetById(id)
	if err != nil {
		return errors.New("配置不存在")
	}
	return this.db.Model(&model).Updates(UpdateData{
		"deleted_at": time.Now(),
		"key":        model.Key + time.Now().Format("2006-01-02 15:04:05"),
	}).Error
}

// 按id查找
func (this *SystemConfigRepositories) GetById(id int) (*models.SysConfig, error) {
	var model = &models.SysConfig{}
	err := this.db.First(model, id).Error
	return model, err
}

// 按 key 查找
func (this *SystemConfigRepositories) GetByKey(name string) (*models.SysConfig) {
	var model = &models.SysConfig{}
	this.db.Where("`key` = ?", name).First(model)
	return model
}

// 单独获取 key value 字端
func (this *SystemConfigRepositories) GetConfigKv() (list []models.SysConfig, err error) {
	err = this.db.Select([]string{"`key`", "`value`"}).Find(&list).Error
	return list, err
}
