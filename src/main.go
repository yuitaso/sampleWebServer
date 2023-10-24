package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/yuitaso/sampleWebServer/src/entity"
	itemHandler "github.com/yuitaso/sampleWebServer/src/handler/item"
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
		authorized.POST("item/create", itemHandler.Create)
	}

	r.POST("/user/create", userHandler.Create)

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

// security要件がないので一旦これで。。
const authEmailKey = "X-Email"
const authPasswordKey = "X-Pass"

func authRequired(c *gin.Context) { // TODO いい感じの置き場にGO
	email := c.Request.Header.Get(authEmailKey)
	password := c.Request.Header.Get(authPasswordKey)

	user, err := userManager.VerifyAndGetUser(email, password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized."})
		return
	}

	c.Set(entity.CtxAuthUserKey, user)
	c.Next()
}
