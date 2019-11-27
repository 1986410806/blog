package repositories

import (
	"blog/app/models"
	"github.com/jinzhu/gorm"
)

type ArticleTagRepository struct {
	db *gorm.DB
}

func NewArticleTagRepository(db *gorm.DB) *ArticleTagRepository {
	return &ArticleTagRepository{
		db: db,
	}
}

// 获取文章标签列表
func (this *ArticleTagRepository) List() (list []models.ArticleTag, err error) {
	err = this.db.Find(&list).Error
	return list, err
}

// 创建文章标签
func (this *ArticleTagRepository) Create(articleTag *models.ArticleTag) (*models.ArticleTag, error) {
	err := this.db.Create(&articleTag).Error
	return articleTag, err
}

// 使用 model 编辑文章标签
func (this *ArticleTagRepository) UpdateByModel(tag *models.ArticleTag, data UpdateData) (err error) {
	err = this.db.Model(&tag).Updates(data).Error
	return err
}

// 批量插入
func (this *ArticleTagRepository) BatchCreate(c uint, tagIds []uint) {
	for _, tagId := range tagIds {
		this.Create(&models.ArticleTag{
			ArticleId: int64(tagId),
			TagId:     int64(tagId),
			Status:    models.ArticleTagStatusOk,
		})
	}
}
