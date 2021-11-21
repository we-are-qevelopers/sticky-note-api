package gin

import (
	"github.com/gin-gonic/gin"

	"sticky-note-api/interfaceAdaptor/gin/controllers"
)

type Routing struct {
	UserController *controllers.UsersController
	AuthController *controllers.AuthController
	Gin            *gin.Engine
	Port           string
}

type RouterParam struct {
	UserController *controllers.UsersController
	AuthController *controllers.AuthController
}

func NewRouting(param RouterParam, GinEngine *gin.Engine, Port string) *Routing {
	r := &Routing{
		AuthController: param.AuthController,
		UserController: param.UserController,
		Gin:            GinEngine,
		Port:           Port,
	}
	// corsの設定をginにセット
	r.Gin.Use(CORSConfig())
	r.setRouting()
	return r
}

// ルーティングを定義
func (r *Routing) setRouting() {
	// 認証が必要なエンドポイント（ログインチェックを行う）
	authGroup := r.Gin.Group("/", LoginCheck())
	{
		authGroup.GET("/users/:id", func(c *gin.Context) { r.UserController.View(c) })
	}
	r.Gin.GET("/hello", func(c *gin.Context) { r.UserController.View(c) })
	r.Gin.POST("/signup", func(c *gin.Context) { r.AuthController.Signup(c) })

}

// サーバー実行
func (r *Routing) Run() {
	r.Gin.Run(r.Port)
}
