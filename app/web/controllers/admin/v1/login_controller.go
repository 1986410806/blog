package v1

import (
	"blog/app/responses"
	"blog/app/services"
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple"
)

type LoginController struct {
	Ctx         iris.Context
	UserService services.UserService
	UserToken   *services.UserTokenService
}

func NewLoginController() *LoginController {
	return &LoginController{
		UserService: services.NewUserService(),
		UserToken:   services.NewUserTokenService(),
	}
}

func (this LoginController) PostLogin() *simple.JsonResult {
	var (
		username = this.Ctx.PostValueTrim("username")
		password = this.Ctx.PostValueTrim("password")
		ref      = this.Ctx.FormValue("ref")
	)
	user, err := this.UserService.SignIn(username, password)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	token, err := this.UserToken.MakeToken(user.ID)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}

	return simple.JsonData(responses.UserTokenResponse.UserToken(user, token, ref))
}

func (this *LoginController) AnyLoginOut() *simple.JsonResult {
	return simple.JsonData("a")
}
