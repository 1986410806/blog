package controllers

import (
	"blog/app/repositories"
	"blog/app/responses"
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple"
)

type TestController struct {
	Ctx            iris.Context
	UserRepository *repositories.UserRepository
}

func NewTestController() *TestController {
	return &TestController{
		UserRepository: repositories.NewUserRepository(),
	}
}

// 用户列表接口
// 获取所有合法用户
func (this TestController) Any() *simple.JsonResult {
	list := responses.UserResponse.List(this.UserRepository.List())
	return simple.JsonData(list)
}
