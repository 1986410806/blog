package repositories

import (
	"blog/app/common/jwt"
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
	db := this.db.Preload("Category").
		Preload("User").
		Preload("ArticleTag").
		Preload("ArticleTag.Tag").
		Find(&list)
	if db.RecordNotFound() {
		return list, nil
	}
	return list, db.Error
}

// 创建文章
func (this *ArticleRepository) Create(userId uint, title, summary, content string,
	categoryId int, tags []string) (*models.Article, error) {
	// 创建文章
	// 1. 先创建相关标签
	// 2. 创建文章
	// 3. 创建文章标签 （中间表）

	article := &models.Article{
		UserId:      userId,
		CategoryId:  int64(categoryId),
		Title:       title,
		Content:     content,
		Summary:     summary,
		ContentType: models.ContentTypeMarkdown,
		Status:      models.ArticleStatusDraft,
	}
	// 启动事务
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
		err = articleTagRepository.BatchCreate(article.ID, tagIds)

		return err
	})
	return article, err
}

// 编辑文章
func (this *ArticleRepository) UpdateById(articleId uint, title, summary, content string,
	categoryId int, tags []string) (*models.Article, error) {
	// 创建文章
	// 1. 先创建相关标签
	// 2. 创建文章
	// 3. 删除文章下所有旧标签 （中间表）
	// 4. 创建新标签

	var article, err = this.GetById(articleId)
	if err != nil {
		return nil, errors.New("文章不存在")
	}

	if article.UserId != jwt.GetTokenClaim().UserId {
		return nil, errors.New("无权限编辑～")
	}

	// save 赋值
	article.Title = title
	article.Summary = summary
	article.CategoryId = int64(categoryId)
	article.Content = content

	err = database.Tx(this.db, func(tx *gorm.DB) error {
		var (
			articleTagRepository = NewArticleTagRepository(tx) // 文章标签存储库
			tagRepository        = NewTagRepository(tx)        // 标签存储库
		)
		// 插入标签
		tagIds := tagRepository.BatchCreate(tags)
		err := tx.Save(&article).Error
		if err != nil {
			return err
		}
		// 删除所有文章旧标签
		err = articleTagRepository.ArticleTagDel(article.ID)
		if err != nil {
			return err
		}
		// 文章标签中间表插入数据
		err = articleTagRepository.BatchCreate(article.ID, tagIds)

		return err
	})
	return article, err
}

// 使用 model 编辑文章
func (this *ArticleRepository) UpdateByModel(tag *models.Article, data UpdateData) (err error) {
	err = this.db.Model(&tag).Updates(data).Error
	return err
}

// 使用 id 删除文章
func (this *ArticleRepository) DelById(id uint) error {
	var tag, err = this.GetById(id)
	if err != nil {
		return errors.New("文章不存在")
	}
	return this.db.Delete(tag).Error
}

// 按id查找
func (this *ArticleRepository) GetById(id uint) (*models.Article, error) {
	var article = &models.Article{}
	e := this.db.First(article, uint(id)).Error
	return article, e
}

// 按 name 查找
func (this *ArticleRepository) GetByName(name string) *models.Article {
	var article = &models.Article{}
	this.db.Where("name = ?", name).First(article)
	return article
}
