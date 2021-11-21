package main

import (
	"fmt"
	"sticky-note-api/infra/gin"
	"sticky-note-api/injector"

	originalGin "github.com/gin-gonic/gin"
)

const Port = ":8080"

func main() {
	fmt.Println("hello world")
	// 各ドメインごとのコントローラーを定義
	userController := injector.InjectUserController()
	authController := injector.InjectAuthController()
	stickyNoteController := injector.InjectStickyNoteController()

	ginRouterParam := gin.RouterParam{
		UserController:       userController,
		AuthController:       authController,
		StickyNoteController: stickyNoteController,
	}
	ginEngine := originalGin.Default()

	r := gin.NewRouting(ginRouterParam, ginEngine, Port)
	r.Run()
}
