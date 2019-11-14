package repository

import (
	"blog/app/models"
	"blog/database"
)

type UserRepository struct{}

func NewUserRepository() UserRepository {
	return UserRepository{}
}

func (n UserRepository) List() (users []models.User) {
	err := database.DB().Find(&users).Error
	if err != nil {
		panic("select Error")
	}
	return users
}
