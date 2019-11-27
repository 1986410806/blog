package admin

import (
	"blog/app/models"
	"blog/app/web/responses"
)

var UserTokenResponse = newUserTokenResponse()

type userTokenResponse struct{}

func newUserTokenResponse() userTokenResponse {
	return userTokenResponse{}
}

func (this userTokenResponse) UserToken(user *models.User, token, ref string) responses.Result {
	return responses.Result{
		"token": token,
		"ref":   ref,
		"user":  UserResponse.User(user)}
}
