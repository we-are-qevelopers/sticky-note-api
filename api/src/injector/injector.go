package injector

import (
	"github.com/we-are-qevelopers/sticky-note-api/domain/repository"
	"github.com/we-are-qevelopers/sticky-note-api/infra/gorm"
	"github.com/we-are-qevelopers/sticky-note-api/interfaceAdaptor/gin/controllers"
	"github.com/we-are-qevelopers/sticky-note-api/useCase"

	originalGorm "gorm.io/gorm"
)

func InjectDB() originalGorm.DB {
	sqlhandler := gorm.NewSqlHandler()
	return *sqlhandler
}

//// Userドメイン
/*
UserRepository(interface)に実装であるSqlHandler(struct)を渡し生成する。
*/
func InjectUserRepository() repository.UserRepository {
	sqlHandler := InjectDB()
	return gorm.NewUserRepository(&sqlHandler)
}

func InjectUserUsecase() useCase.UserUsecase {
	UserRepo := InjectUserRepository()
	return useCase.NewUserUsecase(UserRepo)
}

// ginのコントローラー(hander)にuseCaseを埋め込むイメージ
func InjectUserController() *controllers.UsersController {
	return controllers.NewUsersController(InjectUserUsecase())
}

//// Authドメイン
func InjectAuthController() *controllers.AuthController {
	sqlhandler := gorm.NewSqlHandler()
	return controllers.NewAuthController(useCase.NewAuthUsecase(gorm.NewAuthRepository(sqlhandler)))
}

//// StickyNoteドメイン
func InjectStickyNoteController() *controllers.StickyNoteController {
	sqlhandler := gorm.NewSqlHandler()
	return controllers.NewStickyNoteController(useCase.NewStickyNoteUsecase(gorm.NewStickyNoteRepository(sqlhandler)))
}
