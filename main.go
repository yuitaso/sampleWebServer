package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	hello := []byte("Hello World!!!")

	c.JSON(200, gin.H{
		"message": hello,
	})
}

type User struct {
	Name     string
	Password string
}

func userHandler(c *gin.Context) {
	user := User{Name: "yui", Password: "pass"}

	c.JSON(200, gin.H{"name": user.Name, "password": user.Password})
}

type RequestData struct {
	Id string `uri:"id" binding:"required"`
}

func identifiedUserHandler(c *gin.Context) {
	var request RequestData
	if err := c.ShouldBindUri(&request); err != nil {
		// error
		c.JSON(500, gin.H{})
	}

	fmt.Println("request: ", request)
	user := User{Name: "yui", Password: "pass"}
	c.JSON(200, gin.H{"id": request.Id, "name": user.Name, "password": user.Password})
}

func main() {
	r := gin.Default()

	r.GET("/ping", helloHandler)
	r.GET("/user", userHandler)
	r.GET("/user/:id", identifiedUserHandler)

	r.Run()
}
