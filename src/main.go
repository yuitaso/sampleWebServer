package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	userHandler "github.com/yuitaso/sampleWebServer/src/handler/user"
	userManager "github.com/yuitaso/sampleWebServer/src/manager/user"
)

func main() {
	r := gin.Default()

	authorized := r.Group("/")
	authorized.Use(authRequired)
	{
		authorized.GET("user/:id", userHandler.GetOneById)
		authorized.GET("user/me", userHandler.GetUserMe)
	}

	r.POST("/user/create", userHandler.Create)
	r.POST("/authenticate", userHandler.Authenticate) // TODO handlerの置き場変える

	internalGroup := r.Group("/internal")
	{
		internalGroup.GET("/ping", healthCheckHandler)
	}

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("起動失敗") // fix me
	}
}

func healthCheckHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

// security要件がないので一旦これで。
const authEmailKey = "X-Email"
const authPasswordKey = "X-Pass"

const ContextAuthUserKey = "authUser"

func authRequired(c *gin.Context) {
	email := c.Request.Header.Get(authEmailKey)
	password := c.Request.Header.Get(authPasswordKey)

	user, err := userManager.VerifyPassword(email, password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized."})
		return
	}

	c.Set(ContextAuthUserKey, user)
	c.Next()
}
