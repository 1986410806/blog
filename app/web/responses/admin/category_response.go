package admin

import (
	"blog/app/models"
	"blog/app/web/responses"
)

type CategoryResponse struct {
}

func (r CategoryResponse) Categorys(categorys []*models.Category) (list responses.Results) {
	for _, category := range categorys {
		list = append(list, r.Category(category))
	}
	return list
}

func (CategoryResponse) Category(category *models.Category) responses.Result {
	if category == nil {
		return responses.Result{}
	}
	return responses.Result{
		"id":          category.ID,
		"name":        category.Name,
		"description": category.Description,
		"created_at":  category.CreatedAt,
		"updated_at":  category.UpdatedAt,
	}
}
