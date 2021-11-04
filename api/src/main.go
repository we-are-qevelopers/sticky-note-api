package main

import (
	"fmt"
	"sticky-note-api/infra/gin"
	"sticky-note-api/injector"
)

func main() {
	fmt.Println("hello world")
	userController := injector.InjectUserController()
	// if err := infra.SetupDB(); err != nil {
	// 	log.Fatalln("failed to setup for db")
	// 	// DB接続できない＝サーバーとして機能しないとみなすとここで落としてもいい？
	// 	// DB接続できない場合にも、ちゃんとレスポンスしてあげる必要がある場合は、落としてはいけない
	// }
	// r := infra.NewRouting(infra.DB)
	r := gin.NewRouting(userController)
	// r := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 	c.String(200, "Hello, World")
	// })

	r.Run()
}
