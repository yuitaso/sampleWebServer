package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuitaso/sampleWebServer/src/auth"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := c.Request.Header.Get("X-Token")
		fmt.Println(t)

		tokenString, err := auth.GenerateToken()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": err.Error()})
		}

		fmt.Println("作成済")

		auth.VelifyToken(tokenString)
		c.Next()
	}
}
