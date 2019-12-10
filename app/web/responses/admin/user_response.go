package admin

import (
	"blog/app/models"
	"blog/app/web/responses"
)

var UserResponse = newUserResponse()

type userResponse struct {
}

func newUserResponse() *userResponse {
	return &userResponse{}
}

func (userResponse userResponse) Users(users []models.User) (list responses.Results) {
	for _, user := range users {
		list = append(list, userResponse.User(&user))
	}
	return list
}

func (userResponse userResponse) User(user *models.User) responses.Result {
	return responses.Result{
		"id":          user.ID,
		"username":    user.Username.String,
		"email":       user.Email.String,
		"nickname":    user.Nickname,
		"avatar":      user.Avatar,
		"status":      user.Status,
		"roles":       user.Roles,
		"type":        user.Type,
		"description": user.Description,
		"created_at":  user.CreatedAt.Format("2006-01-02 15:04:05"),
		"updated_at":  user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
