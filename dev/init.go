package main

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/yuitaso/sampleWebServer/src/env"
	itemManager "github.com/yuitaso/sampleWebServer/src/manager/item"
	pointLogManager "github.com/yuitaso/sampleWebServer/src/manager/pointLog"
	userManager "github.com/yuitaso/sampleWebServer/src/manager/user"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	initDBTables()
}

func initDBTables() {
	// open
	db, err := gorm.Open(sqlite.Open(env.DbName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// migrate
	err = db.Table("user").AutoMigrate(&userManager.UserTable{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Table("item").AutoMigrate(&itemManager.ItemTable{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Table("pointLog").AutoMigrate(&pointLogManager.PointLog{})
	if err != nil {
		log.Fatal(err)
	}

	// closeしなくてよさそ？
}
