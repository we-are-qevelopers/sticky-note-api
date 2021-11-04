package main

import (
	"fmt"
	"sticky-note-api/infra/gin"
	"sticky-note-api/injector"
)

func main() {
	fmt.Println("hello world")
	userController := injector.InjectUserController()
	r := gin.NewRouting(userController)
	r.Run()
}
