package repositories

import (
	"blog/app/models"
	"blog/database"
	"errors"
	"github.com/jinzhu/gorm"
)

type ArticleRepository struct {
	db                   *gorm.DB
	articleTagRepository *ArticleTagRepository
	tagRepository        *TagRepository
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{
		db:                   db,
		articleTagRepository: NewArticleTagRepository(db),
		tagRepository:        NewTagRepository(db),
	}
}

// 获取文章列表
func (this *ArticleRepository) List() (list []models.Article, err error) {
	err = this.db.Find(&list).Error
	return list, err
}

// 创建文章
func (this *ArticleRepository) Create(uid uint, title, summary, content string, category_id int, tags []string) (*models.Article, error) {
	// 创建文章
	// 1. 先创建相关标签
	// 2. 创建文章
	// 3. 创建文章标签 （中间表）

	article := &models.Article{
		UserId:      uid,
		CategoryId:  int64(category_id),
		Title:       title,
		Content:     content,
		Summary:     summary,
		ContentType: models.ContentTypeMarkdown,
		Status:      models.ArticleStatusDraft,
	}

	err := database.Tx(this.db, func(tx *gorm.DB) error {
		var (
			articleTagRepository = NewArticleTagRepository(tx) // 文章标签存储库
			tagRepository        = NewTagRepository(tx)        // 标签存储库
		)
		// 插入标签
		tagIds := tagRepository.BatchCreate(tags)

		err := tx.Create(article).Error

		if err != nil {
			return err
		}
		// 文章标签中间表插入数据
		articleTagRepository.BatchCreate(article.ID, tagIds)

		return nil
	})
	return article, err
}

// 编辑文章
func (this *ArticleRepository) UpdateById(id int, data UpdateData) (*models.Article, error) {
	var tag, err = this.GetById(id)
	if err != nil {
		return nil, errors.New("文章不存在")
	}
	err = this.db.Model(&tag).Updates(data).Error
	return tag, err
}

// 使用 model 编辑文章
func (this *ArticleRepository) UpdateByModel(tag *models.Article, data UpdateData) (err error) {
	err = this.db.Model(&tag).Updates(data).Error
	return err
}

// 使用 id 删除文章
func (this *ArticleRepository) DelById(id int) error {
	var tag, err = this.GetById(id)
	if err != nil {
		return errors.New("文章不存在")
	}
	return this.db.Delete(tag).Error
}

// 按id查找
func (this *ArticleRepository) GetById(id int) (*models.Article, error) {
	var tag = &models.Article{}
	e := this.db.First(tag, uint(id)).Error
	return tag, e
}

// 按 name 查找
func (this *ArticleRepository) GetByName(name string) *models.Article {
	var tag = &models.Article{}
	this.db.Where("name = ?", name).First(tag)
	return tag
}
