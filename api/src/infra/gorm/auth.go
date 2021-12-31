package gorm

import (
	"github.com/we-are-qevelopers/sticky-note-api/domain/model"
	"github.com/we-are-qevelopers/sticky-note-api/domain/repository"

	"gorm.io/gorm"
)

type AuthRepository struct {
	sqlHandler *gorm.DB
}

func NewAuthRepository(sqlHandler *gorm.DB) repository.AuthRepository {
	authRepository := AuthRepository{sqlHandler}
	return &authRepository
}

func (authRepo *AuthRepository) Signup(user *model.SignupUser) (err error) {

	err = nil
	return
}
