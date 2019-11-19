package repositories

import (
	"blog/app/models"
	"blog/database"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: database.DB()}
}

func (this *UserRepository) List() (users []models.User) {
	err := this.db.Find(&users).Error
	if err != nil {
		panic("select Error")
	}
	return users
}

func (this *UserRepository) GetById(id uint) (user *models.User) {
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
