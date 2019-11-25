package responses

import "blog/app/models"


type TagResponse struct {
}

func (r TagResponse) Tags(tags []models.Tag) (list results) {
	for _, tag := range tags {
		list = append(list, r.Tag(&tag))
	}
	return list
}

func (r TagResponse) Tag(model *models.Tag) result {
	return result{
		"id":          model.ID,
		"name":        model.Name,
		"description": model.Description,
		"created_at":  model.CreatedAt.Format("2006-01-02 15:04:05"),
		"updated_at":  model.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
