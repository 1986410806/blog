package repositories

import (
	"blog/app/models"
	"blog/database"
	"github.com/jinzhu/gorm"
	"github.com/mlogclub/simple"
	"time"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: database.DB()}
}

func (this *UserRepository) List(paging *simple.Paging) []*models.User {
	users := make([]*models.User, paging.Limit)
	this.db.Offset(paging.Offset()).Limit(paging.Limit).Find(&users)
	this.db.Model(&models.User{}).Count(&paging.Total)
	return users
}

func (this *UserRepository) GetById(id uint) *models.User {
	var user = &models.User{}
	this.db.First(&user, id)
	return user
}

// 根据 email 查找用户
func (this *UserRepository) GetByEmail(email string) *models.User {
	ret := &models.User{}
	this.db.Take(ret, "email = ?", email)
	return ret
}

// 根据 username 查找用户
func (this *UserRepository) GetByUsername(username string) *models.User {
	ret := &models.User{}
	this.db.Take(ret, "username = ?", username)
	return ret
}

// 数据库更新最新时间
func (this *UserRepository) LastLoginTimeById(user *models.User) {
	this.db.Model(&user).Update("last_login_time", time.Now())
}
