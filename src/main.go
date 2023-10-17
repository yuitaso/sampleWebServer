package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/yuitaso/sampleWebServer/src/entities/user"
	userManager "github.com/yuitaso/sampleWebServer/src/entities/user/manager"
	"strconv"
)

func helloHandler(c *gin.Context) {
	hello := []byte("Hello World!!!")

	c.JSON(200, gin.H{
		"message": hello,
	})
}

func userHandler(c *gin.Context) {
	newUser := user.User{Name: "aaa", Password: "xxx"}

	c.JSON(200, gin.H{"name": newUser.Name, "password": newUser.Password})
}

type RequestData struct {
	Id string `uri:"id" binding:"required"`
}

func identifiedUserHandler(c *gin.Context) {
	var request RequestData
	var id int

	err := c.ShouldBindUri(&request)
	id, err = strconv.Atoi(request.Id)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()}) // いい感じに返すConfがあるはず
	}

	res := userManager.FindById(id)

	c.JSON(200, gin.H{"id": request.Id, "name": res.Name, "password": res.Password})
}

func createUserHandler(c *gin.Context) {
	err := userManager.Create(user.User{Name: "gorm try", Password: "hoge"})
	if err != nil {
		c.JSON(500, gin.H{"message": "できんかった"})
	}

	c.JSON(200, gin.H{"message": "できた"})
}

func main() {
	r := gin.Default()

	r.GET("/ping", helloHandler)
	r.GET("/user", userHandler)
	r.GET("/user/:id", identifiedUserHandler)
	r.GET("/user/create", createUserHandler)
	r.Run()
}
