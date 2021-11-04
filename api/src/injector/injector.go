package injector

import (
	"sticky-note-api/domain/repository"
	"sticky-note-api/infra/gorm"
	"sticky-note-api/interfaceAdaptor/gin/controllers"
	"sticky-note-api/useCase"

	originalGorm "gorm.io/gorm"
)

func InjectDB() originalGorm.DB {
	sqlhandler := gorm.NewSqlHandler()
	return *sqlhandler
}

/*
UserRepository(interface)に実装であるSqlHandler(struct)を渡し生成する。
*/
func InjectUserRepository() repository.UserRepository {
	sqlHandler := InjectDB()
	// var aaa sqlHandler
	return gorm.NewUserRepository(&sqlHandler)
}

func InjectUserUsecase() useCase.UserUsecase {
	UserRepo := InjectUserRepository()
	return useCase.NewUserUsecase(UserRepo)
}

func InjectUserController() *controllers.UsersController {
	return controllers.NewUsersController(InjectUserUsecase())
}
