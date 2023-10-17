package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/yuitaso/sampleWebServer/env"
	userManager "github.com/yuitaso/sampleWebServer/src/entities/user/manager"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {
	initDBTables()
}

func initDBTables() {
	db, err := gorm.Open(sqlite.Open(env.DbName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Table("user").AutoMigrate(&userManager.UserModel{})
	if err != nil {
		log.Fatal(err)
	}
}
