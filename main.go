package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	userEntity "github.com/yuitaso/sampleWebServer/entities/user"
	"github.com/yuitaso/sampleWebServer/env"
	"log"
)

func helloHandler(c *gin.Context) {
	hello := []byte("Hello World!!!")

	c.JSON(200, gin.H{
		"message": hello,
	})
}

func userHandler(c *gin.Context) {
	newUser := userEntity.User{Name: "yui", Password: "pass"}

	c.JSON(200, gin.H{"name": newUser.Name, "password": newUser.Password})
}

type RequestData struct {
	Id string `uri:"id" binding:"required"`
}

func identifiedUserHandler(c *gin.Context) {
	var request RequestData
	if err := c.ShouldBindUri(&request); err != nil {
		// error
		c.JSON(500, gin.H{}) // いい感じに返すConfがあるはず
	}
	fmt.Println("おけ??")
	fmt.Println(request.Id)

	// open db
	db, err := sql.Open("sqlite3", env.DbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// create statement
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("select name, pass from user where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// exec
	var data_user userEntity.User
	err = stmt.QueryRow(1).Scan(&data_user.Name, &data_user.Password)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{"id": request.Id, "name": data_user.Name, "password": data_user.Password})
}

func withManagerHandler(c *gin.Context) {
	var request RequestData
	if err := c.ShouldBindUri(&request); err != nil {
		// error
		c.JSON(500, gin.H{}) // いい感じに返すConfがあるはず
	}
	fmt.Println("おけ??")
	fmt.Println(request.Id)

	// open db
	db, err := sql.Open("sqlite3", env.DbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// create statement
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("select name, pass from user where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// exec
	var data_user userEntity.User
	err = stmt.QueryRow(1).Scan(&data_user.Name, &data_user.Password)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{"id": request.Id, "name": data_user.Name, "password": data_user.Password})
}

func dbConnHandler(c *gin.Context) {
	fmt.Println("createするよ〜")

	// open db
	db, err := sql.Open("sqlite3", env.DbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// create statement
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into user(name, pass) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// exec
	_, err = stmt.Exec("yui", "pass")
	if err != nil {
		log.Fatal(err)
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{"message": "でけた"})
}

func main() {
	r := gin.Default()

	r.GET("/ping", helloHandler)
	r.GET("/user", userHandler)
	r.GET("/user/:id", identifiedUserHandler)
	r.GET("/user/create", dbConnHandler)
	r.GET("/user/list", withManagerHandler)
	r.Run()
}
