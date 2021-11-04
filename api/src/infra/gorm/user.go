package gorm

import (
	"sticky-note-api/domain/model"
	"sticky-note-api/domain/repository"

	"gorm.io/gorm"
)

type UserRepository struct {
	sqlHandler *gorm.DB
}

func NewUserRepository(sqlHandler *gorm.DB) repository.UserRepository {
	userRepository := UserRepository{sqlHandler}
	return &userRepository
}

func (userRepo *UserRepository) FindAll() (users []*model.User, err error) {
	users = []*model.User{
		{ID: 1, Name: "testuser"},
	}
	err = nil
	return
}

func (userRepo *UserRepository) Create(todo *model.User) (*model.User, error) {
	user := model.User{ID: 1, Name: "testuser"}

	return &user, nil
}
