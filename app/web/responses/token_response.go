package responses

import (
	"blog/app/models"
)

var UserTokenResponse = newUserTokenResponse()

type userTokenResponse struct{}

func newUserTokenResponse() userTokenResponse {
	return userTokenResponse{}
}

func (this userTokenResponse) UserToken(user *models.User, token, ref string) result {
	return result{
		"token": token,
		"ref":   ref,
		"user":  UserResponse.User(user)}
}