package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/toma-photo/internal/models/common/response"
	"github.com/toma-photo/internal/utils"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		if token == "" {
			response.Unauthorized(c)
			c.Abort()
			return
		}
		j := utils.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			response.Unauthorized(c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
