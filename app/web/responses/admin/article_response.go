package admin

import (
	"blog/app/models"
	"blog/app/web/responses"
)

type ArticleResponse struct {
	CategoryResponse CategoryResponse
	TagResponse      TagResponse
}

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

func (res ArticleResponse) List(articles []*models.Article) (results responses.Results) {
	for _, article := range articles {
		result := responses.Result{
			"id":           article.ID,
			"title":        article.Title,
			"summary":      article.Summary,
			"status":       article.Status,
			"content_type": article.ContentType,
			"content":      article.Content,
			"tag":          res.ArticleTags(article.ArticleTag),
			"category":     res.CategoryResponse.Category(article.Category),
		}
		results = append(results, result)
	}
	return results
}

//
func (res ArticleResponse) ArticleTags(articleTags []*models.ArticleTag) (results responses.Results) {
	for _, articleTag := range articleTags {
		tag := res.TagResponse.Tag(articleTag.Tag)
		results = append(results, tag)
	}
	return results
}
