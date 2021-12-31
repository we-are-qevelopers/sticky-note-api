package repository

import (
	"github.com/we-are-qevelopers/sticky-note-api/domain/model"
)

type AuthRepository interface {
	Signup(user *model.SignupUser) (err error)
}
