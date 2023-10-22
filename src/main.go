package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	userHandler "github.com/yuitaso/sampleWebServer/src/handler/user"
	middleware "github.com/yuitaso/sampleWebServer/src/middleWare"
)

func main() {
	r := gin.Default()

	authorized := r.Group("/")
	authorized.Use(middleware.AuthRequired())
	{
		authorized.GET("user/:id", userHandler.GetOneById)
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
