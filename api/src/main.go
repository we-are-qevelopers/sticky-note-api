package main

import (
	"fmt"
	"sticky-note-api/infra/gin"
	"sticky-note-api/injector"
)

func main() {
	fmt.Println("hello world")
	userController := injector.InjectUserController()
	authController := injector.InjectAuthController()

	ginRouterParam := gin.RouterParam{
		UserController: userController,
		AuthController: authController,
	}

	r := gin.NewRouting(ginRouterParam)
	r.Run()
}
