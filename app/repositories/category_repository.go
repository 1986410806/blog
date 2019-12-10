package repositories

import (
	"blog/app/models"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/mlogclub/simple"
	"time"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db}
}

// 获取栏目列表
func (this *CategoryRepository) CategoryList(paging *simple.Paging) []*models.Category {
	list := make([]*models.Category, paging.Limit)
	this.db.Offset(paging.Offset()).Limit(paging.Limit).Find(&list)
	this.db.Model(&models.Category{}).Count(&paging.Total)
	return list
}

// 创建栏目
func (this *CategoryRepository) Create(Category *models.Category) (*models.Category, error) {
	if tmp := this.GetByName(Category.Name); tmp.Name == Category.Name {
		return tmp, nil
	}
	err := this.db.Create(&Category).Error
	return Category, err
}

// 编辑栏目
func (this *CategoryRepository) UpdateById(id int, data UpdateData) (*models.Category, error) {
	var Category, err = this.GetById(id)
	if err != nil {
		return nil, errors.New("栏目不存在")
	}
	err = this.db.Model(&Category).Updates(data).Error
	return Category, err
}

// 使用 model 编辑栏目
func (this *CategoryRepository) UpdateByModel(Category *models.Category, data UpdateData) (err error) {
	err = this.db.Model(&Category).Updates(data).Error
	return err
}

// 使用id 删除栏目
func (this *CategoryRepository) DelById(id int) error {
	var category, err = this.GetById(id)
	if err != nil {
		return errors.New("栏目不存在")
	}
	// 为避免唯一索引冲突 软删除的时候给唯一列拼接上 当前时间
	// 也可以数据库上使用和软删除字端组合唯一列
	return this.db.Model(&category).Updates(UpdateData{
		"deleted_at": time.Now(),
		"name":       category.Name + time.Now().Format("2006-01-02 15:04:05"),
	}).Error
}

// 按id查找
func (this *CategoryRepository) GetById(id int) (*models.Category, error) {
	var Category = &models.Category{}
	e := this.db.First(Category, uint(id)).Error
	return Category, e
}

// 按 name 查找
func (this *CategoryRepository) GetByName(name string) *models.Category {
	var Category = &models.Category{}
	this.db.Where("name = ?", name).First(Category)
	return Category
}
