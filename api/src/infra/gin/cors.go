package gin

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSConfig() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowMethods: []string{
			"POST",
		},
		AllowOrigins: []string{
			"*",
		},
		AllowHeaders: []string{
			"Content-Type",
		},
	})
}
