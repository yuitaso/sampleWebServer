package main

import (
	"flag"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/yuitaso/sampleWebServer/src/env"
	contractManager "github.com/yuitaso/sampleWebServer/src/manager/contract"
	itemManager "github.com/yuitaso/sampleWebServer/src/manager/item"
	pointLogManager "github.com/yuitaso/sampleWebServer/src/manager/pointLog"
	userManager "github.com/yuitaso/sampleWebServer/src/manager/user"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	var (
		e = flag.String("e", "dev", "string flag")
	)
	flag.Parse()
	fmt.Println("Execute with env: ", *e)

	env.SetEnv(*e)

	fmt.Println("RUNS ON >>> ------------")
	fmt.Println("env = ", env.Env.Env)
	fmt.Println("db = ", env.Env.DbName)
	fmt.Println("------------------------")
}

func main() {
	initDBTables()
}

func initDBTables() {
	// open
	db, err := gorm.Open(sqlite.Open(env.Env.DbName), &gorm.Config{})
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
	err = db.Table("pointLog").AutoMigrate(&pointLogManager.PointLogTable{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Table("contract").AutoMigrate(&contractManager.ContractTable{})
	if err != nil {
		log.Fatal(err)
	}

	// closeしなくてよさそ？
}
