package controllers

import (
	"blog/app/repository"
	"blog/app/responses"
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple"
)

type TestController struct {
	Ctx iris.Context
	//Service service.BookService
	UserRepository repository.UserRepository
	UserResponse   responses.UserResponse
}

func NewTestController() *TestController {
	return &TestController{
		UserRepository: repository.NewUserRepository(),
		UserResponse:   responses.NewUserResponse()}
}

// 用户列表接口
// 获取所有合法用户
func (this *TestController) Any() *simple.JsonResult {
	list := this.UserResponse.List(this.UserRepository.List())
	return simple.JsonData(list)
}
