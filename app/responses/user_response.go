package responses

import "blog/app/models"

type UserResponse struct {
}

func NewUserResponse() UserResponse {
	return UserResponse{}
}

func (userResponse UserResponse) List(users []models.User) (list []map[string]interface{}) {
	for _, user := range users {
		list = append(list, userResponse.Item(&user))
	}
	return list
}

func (userResponse UserResponse) Item(user *models.User) map[string]interface{} {
	return map[string]interface{}{
		"ID":          user.ID,
		"username":    user.Username.String,
		"email":       user.Email.String,
		"nickname":    user.Nickname,
		"avatar":      user.Avatar,
		"status":      user.Status,
		"roles":       user.Roles,
		"type":        user.Type,
		"description": user.Description,
		"CreatedAt":   user.CreatedAt.Format("2006-01-02 15:04:05"),
		"UpdatedAt":   user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
