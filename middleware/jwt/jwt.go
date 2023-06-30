package jwt

import (
	"TreeHole/treehole_backend/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		var errMsg string

		code = 200
		token := c.Query("token")
		if token == "" {
			code = 400
			errMsg = "Token缺失"
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = 0
				errMsg = "Token验证失败"
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = 0
				errMsg = "Token已超时"
			}
		}

		if code != 200 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  errMsg,
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
