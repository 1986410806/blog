package v1

import (
	"blog/app/common/jwt"
	"blog/app/repositories"
	"blog/app/web/responses/admin"
	"blog/app/web/services"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple"
	"github.com/sirupsen/logrus"
)

func Login(ctx iris.Context) {

	userService := services.NewUserService()
	userRepository := repositories.NewUserRepository()

	var (
		username = ctx.PostValueTrim("username")
		password = ctx.PostValueTrim("password")
		ref      = ctx.FormValue("ref")
	)
	user, err := userService.SignIn(username, password)
	if err != nil {
		ctx.JSON(simple.JsonErrorMsg(err.Error()))
	}
	token, err := jwt.MakeToken(user.ID, user.Username.String, user.Email.String, user.Roles)
	if err != nil {
		ctx.JSON(simple.JsonErrorMsg(err.Error()))
	}
	userRepository.LastLoginTimeById(user)
	// 用户登录 记录日志
	logrus.Info(fmt.Sprintf("用户[%s]登录了管理后台;id[%d]", user.Username.String, user.ID))
	ctx.JSON(
		simple.JsonData(
			admin.UserTokenResponse.UserToken(
				user, token, ref)))
}

func Logout(ctx iris.Context) {
	var token = jwt.GetTokenClaim()
	logrus.Info(fmt.Sprintf("用户[%s]退出登录;id[%d]", token.UserName, token.UserId))
	ctx.JSON(
		simple.JsonData("res:ok"))
}
