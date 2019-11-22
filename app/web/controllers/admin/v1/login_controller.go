package v1

import (
	"blog/app/common/jwt"
	"blog/app/repositories"
	"blog/app/services"
	"blog/app/web/responses"
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple"
)

type LoginController struct {
	Ctx            iris.Context
	UserService    services.UserService
	UserRepository *repositories.UserRepository
}

func NewLoginController() *LoginController {
	return &LoginController{
		UserService:    services.NewUserService(),
		UserRepository: repositories.NewUserRepository(),
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
	token, err := jwt.MakeToken(user)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	this.UserRepository.LastLoginTimeById(user)

	return simple.JsonData(responses.UserTokenResponse.UserToken(user, token, ref))
}

func (this *LoginController) AnyLoginOut() *simple.JsonResult {
	return simple.JsonData("res:ok")
}
