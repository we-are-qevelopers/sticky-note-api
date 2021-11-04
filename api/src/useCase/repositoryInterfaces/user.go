package repositoryinterfaces

import (
	"sticky-note-api/domain"

	"gorm.io/gorm"
)

/*
永続化させるための関数のインターフェイスを定義
*/
type IUsersRepository interface {
	Store(domain.User) (int, error)
	FindByID(db *gorm.DB, id int) (name string, err error)
}
