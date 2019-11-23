package v1

import (
	"blog/app/common/jwt"
	"blog/app/repositories"
	"blog/app/web/responses"
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple"
)

type UserController struct {
	Ctx            iris.Context
	UserRepository *repositories.UserRepository
}

func NewUserController() *UserController {
	return &UserController{
		UserRepository: repositories.NewUserRepository(),
	}
}

/**
 * 用户信息
 * @return blog/app/repositories/UserResponse/User
 */
func (c *UserController) Get() *simple.JsonResult {

	TokenClaim := jwt.GetTokenClaim(c.Ctx)

	user := c.UserRepository.GetById(TokenClaim.UserId)

	return simple.JsonData(responses.UserResponse.User(user))
}

/***
 * 用户列表
 * @param
 * @return *simple.JsonResult
 */
func (c *UserController) GetList() *simple.JsonResult {

	TokenClaim := jwt.GetTokenClaim(c.Ctx)

	user := c.UserRepository.GetById(TokenClaim.UserId)

	return simple.JsonData(responses.UserResponse.User(user))
}
