package services

import (
	"blog/app/common"
	"blog/app/models"
	"blog/app/repositories"
	"errors"
	"fmt"
	"github.com/mlogclub/simple"
)

type UserService interface {
	SignIn(username, password string) (*models.User, error)
}

func NewUserService() UserService {
	return &userService{
		userRepository: repositories.NewUserRepository(),
	}
}

type userService struct {
	userRepository *repositories.UserRepository
}

// 登录
func (this *userService) SignIn(username, password string) (*models.User, error) {
	if len(username) == 0 {
		return nil, errors.New("用户名/邮箱不能为空")
	}
	if len(password) == 0 {
		return nil, errors.New("密码不能为空")
	}
	var user *models.User = nil
	if err := common.IsValidateEmail(username); err == nil { // 如果用户输入的是邮箱
		user = this.userRepository.GetByEmail(username)
	} else {
		fmt.Println(this)
		user = this.userRepository.GetByUsername(username)
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}
	if !simple.ValidatePassword(user.Password, password) {
		return nil, errors.New("密码错误")
	}
	return user, nil
}
