package gin

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"sticky-note-api/interfaceAdaptor/gin/controllers"
)

type Routing struct {
	UserController *controllers.UsersController
	Gin            *gin.Engine
	Port           string
}

func NewRouting(userController *controllers.UsersController) *Routing {
	r := &Routing{
		UserController: userController,
		Gin:            gin.Default(),
		Port:           ":8080",
	}
	r.setRouting()
	return r
}

func (r *Routing) setRouting() {
	r.Gin.GET("/", func(c *gin.Context) { r.UserController.View(c) })
	r.Gin.GET("/users/:id", func(c *gin.Context) { r.UserController.View(c) })
}

func (r *Routing) Run() {
	fmt.Println(9999)
	r.Gin.Run(r.Port)
}
