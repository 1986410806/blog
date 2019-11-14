package admin

import (
	"blog/app/repository"
	"blog/app/responses"
	"github.com/kataras/iris/v12"
)

type LoginController struct {
	Ctx iris.Context
	//Service service.BookService
	UserRepository repository.UserRepository
	UserResponse   responses.UserResponse
}

func NewLoginController() *LoginController {
	return &LoginController{}
}
