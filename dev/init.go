package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/yuitaso/sampleWebServer/src/env"
	userManager "github.com/yuitaso/sampleWebServer/src/manager/user"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"flag"
	"fmt"
)

func init() {
    flag.Parse()
    fmt.Println(flag.Args())
}

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
}
