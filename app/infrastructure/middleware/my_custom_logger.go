package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func MyCustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("-------MyCustomLogger Start------")
		c.Next()
		fmt.Println("-------MyCustomLogger End------")
	}
}
