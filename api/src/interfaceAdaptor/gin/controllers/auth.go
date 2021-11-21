package controllers

import (
	"fmt"
	"sticky-note-api/useCase"
)

type AuthController struct {
	authUseCase useCase.AuthUsecase
}

func NewAuthController(authUseCase useCase.AuthUsecase) *AuthController {
	return &AuthController{authUseCase: authUseCase}
}

type ReqBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (controller *AuthController) Signup(c Context) {

	var reqBody ReqBody
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		fmt.Println(err)
	}
	fmt.Println(reqBody)

	// users, err := controller.authUseCase.Signup()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(users)

	c.JSON(200, "Sign up")
}
