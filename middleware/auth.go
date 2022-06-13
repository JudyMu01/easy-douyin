package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func JWY() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = 200
		token := c.Query("token")

		if token == "" {
			code = 401
		} else {
			claims, err := ParseToken(token)
			if err != nil {
				code = 402
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = 403
			}
		}

		//如果token验证不通过，直接终止程序，c.Abort()
		if code != 200 {
			// 返回错误信息
			c.JSON(http.StatusUnauthorized, http.Response{StatusCode: code})

			return
		}
		c.Next()
	}
}
