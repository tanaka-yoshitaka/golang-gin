package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/44taka/golang-gin/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.GetHeader("Authorization")
		if clientToken == "" {
			c.JSON(http.StatusForbidden, "No Authorization header provided")
			c.Abort()
		}

		extractedToken := strings.Split(clientToken, "Bearer ")
		fmt.Println(extractedToken[0])
		fmt.Println(extractedToken[1])

		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			c.JSON(400, "Incorrect Format of Authorization Token")
			c.Abort()
			return
		}

		// TODO::シークレットキーなどは設定ファイルから読み込ませる
		jwtWrapper := utils.JwtWrapper{
			SecretKey: "verysecretkey",
			Issuer:    "AuthService",
		}

		claims, err := jwtWrapper.ValidateToken(clientToken)
		if err != nil {
			c.JSON(401, err.Error())
			c.Abort()
			return
		}

		c.Set("id", claims.UserId)
		c.Next()
	}
}
