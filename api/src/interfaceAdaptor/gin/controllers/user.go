package controllers

import (
	"fmt"
	"github.com/we-are-qevelopers/sticky-note-api/useCase"
)

type UsersController struct {
	userUseCase useCase.UserUsecase
}

func NewUsersController(userUseCase useCase.UserUsecase) *UsersController {
	return &UsersController{userUseCase: userUseCase}
}

/*
コントローラーは、
・useCaseのinputに向けてデータを変換
・useCaseのoutputのデータを変換
する
*/
func (controller *UsersController) View(c Context) {

	users, err := controller.userUseCase.View()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(users)

	c.JSON(200, "hello world")
}
