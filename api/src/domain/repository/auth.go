package repository

import (
	"sticky-note-api/domain/model"
)

type AuthRepository interface {
	Signup(user *model.SignupUser) (err error)
}
