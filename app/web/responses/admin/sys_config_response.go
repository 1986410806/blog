package admin

import (
	"blog/app/models"
	"blog/app/web/responses"
)

type SystemConfigResponse struct {
}

func (r SystemConfigResponse) List(models []models.SysConfig) (list responses.Results) {
	for _, model := range models {
		list = append(list, r.Item(&model))
	}
	return list
}

func (r SystemConfigResponse) Item(model *models.SysConfig) responses.Result {
	return responses.Result{
		"id":          model.ID,
		"key":         model.Key,
		"value":       model.Value,
		"name":        model.Name,
		"description": model.Description,
		"created_at":  model.CreatedAt.Format("2006-01-02 15:04:05"),
		"updated_at":  model.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
