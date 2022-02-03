package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.GetHeader("X-Request-ID")
		if requestId == "" {
			u, _ := uuid.NewRandom()
			requestId = u.String()
		}
		c.Set("requestId", requestId)
		c.Next()
	}
}
