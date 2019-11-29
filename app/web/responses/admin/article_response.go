package admin

import (
	"blog/app/models"
	"blog/app/web/responses"
)

type ArticleResponse struct{}

// 新增文章
func (ArticleResponse) Create(article *models.Article) responses.Result {
	return responses.Result{
		"id":           article.ID,
		"title":        article.Title,
		"summary":      article.Summary,
		"status":       article.Status,
		"content_type": article.ContentType,
		"content":      article.Content,
	}
}
