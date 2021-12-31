package useCase

import (
	"github.com/we-are-qevelopers/sticky-note-api/domain/model"
	"github.com/we-are-qevelopers/sticky-note-api/domain/repository"
)

type AuthUsecase interface {
	Signup(signupUser *model.SignupUser) (err error)
}

type authUsecase struct {
	authRepo repository.AuthRepository
}

func NewAuthUsecase(authRepository repository.AuthRepository) AuthUsecase {
	newAuthUsecase := authUsecase{authRepo: authRepository}

	return &newAuthUsecase
}

func (usecase *authUsecase) Signup(user *model.SignupUser) (err error) {
	err = usecase.authRepo.Signup(user)
	return
}
