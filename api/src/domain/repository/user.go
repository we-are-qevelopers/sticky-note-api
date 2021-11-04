package repository

import (
	"sticky-note-api/domain/model"
)

// domain/repositoryには、永続化させるための関数のinterfaceを定義
type UserRepository interface {
	FindAll() (users []*model.User, err error)
	Create(user *model.User) (*model.User, error)
}
