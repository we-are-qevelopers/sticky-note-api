package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 一旦httpリクエストにuserIDがあるかないかで、ログインチェックを行うようにしている
		userID := c.Request.Header.Get("userID")

		if userID == "" {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		// 実際は、ここで認証チェックを行う

		c.Set("userID", userID)
	}
}
