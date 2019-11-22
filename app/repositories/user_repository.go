package repositories

import (
	"blog/app/models"
	"blog/database"
	"blog/database/cache"
	cache2 "github.com/go-redis/cache/v7"
	"github.com/jinzhu/gorm"
	"time"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: database.DB()}
}

func (this *UserRepository) List() (users []models.User) {
	var cached = cache.GetCache()
	err := cached.Once(&cache2.Item{
		Key:        "user:list",
		Object:     &users,
		Expiration: 100 * time.Second,
		Func: func() (i interface{}, e error) {
			err := this.db.Find(&users).Error
			return users, err
		},
	})
	if err != nil {
		panic(err)
	}
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
