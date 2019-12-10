package controllers

import (
	"blog/app/repositories"
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

	var page = &simple.Paging{
		Page:  simple.FormValueIntDefault(this.Ctx, "page", 1),
		Limit: simple.FormValueIntDefault(this.Ctx, "limit", 10),
		Total: 0,
	}

	list := this.UserRepository.List(page)

	return simple.JsonData(
		simple.PageResult{
			Page:    page,
			Results: list,
		})
}
