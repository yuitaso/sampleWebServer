package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuitaso/sampleWebServer/src/auth"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("処理前")

		fmt.Println(c.Request.Header.Get("X-Token"))

		token, err := auth.GenerateToken()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": err.Error()})
		}

		fmt.Println("認証おっけー", token)
		c.Next()
	}
}
