package v1

import (
	"blog/app/repositories"
	"blog/app/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12"
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
	Claims := c.Ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	id := (Claims["id"]).(float64)
	user := c.UserRepository.GetById(uint(id))
	return simple.JsonData(user)
}
