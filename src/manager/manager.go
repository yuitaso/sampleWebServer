package manager

import (
	"log"

	"github.com/yuitaso/sampleWebServer/src/env"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func OpenDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open(env.DbName), &gorm.Config{})
	if err != nil {
		log.Fatal("cannnot connect to database.")
	}
}