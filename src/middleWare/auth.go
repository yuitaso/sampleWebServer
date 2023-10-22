package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuitaso/sampleWebServer/src/auth"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("X-Token")
		fmt.Println("とーくんきたよ", tokenString)

		token, err := auth.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Unauthrized."})
		}

		fmt.Println(token.Claims.(*auth.AuthClaims).Uuid)
		c.Next()
	}
}
