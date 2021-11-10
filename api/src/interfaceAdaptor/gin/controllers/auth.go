package controllers

import (
	"sticky-note-api/useCase"
)

type AuthController struct {
	authUseCase useCase.AuthUsecase
}

func NewAuthController(authUseCase useCase.AuthUsecase) *AuthController {
	return &AuthController{authUseCase: authUseCase}
}

func (controller *AuthController) Signup(c Context) {

	// users, err := controller.authUseCase.Signup()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(users)

	c.JSON(200, "hello world")
}
