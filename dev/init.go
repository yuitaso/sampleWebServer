package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/yuitaso/sampleWebServer/env"
	"log"
)

func main() {
	initDBTables()
}

func initDBTables() {
	db, err := sql.Open("sqlite3", env.DbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// create User table
	sqlStmt := `
	create table user (id integer not null primary key, name text, pass text);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
	}

}
