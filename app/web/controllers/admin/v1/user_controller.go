package v1

import (
	"blog/app/repositories"
	"blog/app/services"
	"github.com/kataras/iris"
	"github.com/mlogclub/simple"
)

type UserController struct {
	Ctx            iris.Context
	UserRepository *repositories.UserRepository
	UserToken      *services.UserTokenService
}

func NewUserController() *UserController {
	return &UserController{
		UserRepository: repositories.NewUserRepository(),
		UserToken:      services.NewUserTokenService(),
	}
}

func (c *UserController) Get() *simple.JsonResult {
	//var id  = c.Ctx.FormValue('')
	//c.UserRepository.GetById(id)
}
