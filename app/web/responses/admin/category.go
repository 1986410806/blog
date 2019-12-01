package admin

import (
	"blog/app/models"
	"blog/app/web/responses"
)

type CategoryResponse struct{}

func (CategoryResponse) Category(category *models.Category) responses.Result {
	return responses.Result{
		"id":          category.ID,
		"name":        category.Name,
		"description": category.Description,
		"created_at":  category.CreatedAt,
		"updated_at":  category.UpdatedAt,
	}
}
