package responses

import "blog/app/models"

var UserResponse = newUserResponse()

type userResponse struct {
}

func newUserResponse() *userResponse {
	return &userResponse{}
}

func (userResponse userResponse) List(users []models.User) (list Results) {
	for _, user := range users {
		list = append(list, userResponse.Item(&user))
	}
	return list
}

func (userResponse userResponse) Item(user *models.User) Result {
	return Result{
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
