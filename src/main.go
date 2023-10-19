package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	userHandler "github.com/yuitaso/sampleWebServer/src/handler/user"
	"log"
)

func main() {
	r := gin.Default()

	userGroup := r.Group("/user")
	{
		userGroup.POST("/create", userHandler.Create)
		userGroup.GET("/:id", userHandler.GetOneById)
		userGroup.POST("/authenticate", userHandler.Authenticate)
	}

	internalGroup := r.Group("/internal")
	{
		internalGroup.GET("/ping", healthCheckHandler)
	}

	err := r.Run()
	if err != nil {
		log.Fatal("起動失敗") // fix me
	}
}

func healthCheckHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
