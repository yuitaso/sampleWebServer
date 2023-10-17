package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	userHandler "github.com/yuitaso/sampleWebServer/src/handler/user"
	"log"
)

func main() {
	r := gin.Default()

	r.GET("/ping", healthCheckHandler)

	userGroup := r.Group("/user")
	{
		userGroup.GET("/create", userHandler.Create)
		userGroup.GET("/:id", userHandler.GetOneById)
	}

	err := r.Run()
	if err != nil {
		log.Fatal("起動失敗")
	}
}

func healthCheckHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
