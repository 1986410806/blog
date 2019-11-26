package responses

import "blog/app/models"

type SystemConfigResponse struct {
}

func (r SystemConfigResponse) List(models []models.SysConfig) (list results) {
	for _, model := range models {
		list = append(list, r.Item(&model))
	}
	return list
}

func (r SystemConfigResponse) Item(model *models.SysConfig) result {
	return result{
		"id":          model.ID,
		"key":         model.Key,
		"value":       model.Value,
		"name":        model.Name,
		"description": model.Description,
		"created_at":  model.CreatedAt.Format("2006-01-02 15:04:05"),
		"updated_at":  model.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
