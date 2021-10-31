package main

import (
	"fmt"
	"log"
	"sticky-note-api/infra"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world")
	if err := infra.SetupDB(); err != nil {
		log.Fatalln("failed to setup for db")
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World")
	})
	r.Run()

}
