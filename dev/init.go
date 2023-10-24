package main

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/yuitaso/sampleWebServer/src/env"
	itemManager "github.com/yuitaso/sampleWebServer/src/manager/item"
	userManager "github.com/yuitaso/sampleWebServer/src/manager/user"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	initDBTables()
}

func initDBTables() {
	db, err := gorm.Open(sqlite.Open(env.DbName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Table("user").AutoMigrate(&userManager.UserTable{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Table("item").AutoMigrate(&itemManager.ItemTable{})
	if err != nil {
		log.Fatal(err)
	}
}
