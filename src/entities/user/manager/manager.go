package manager

import (
	"database/sql"
	"github.com/yuitaso/sampleWebServer/env"
	"github.com/yuitaso/sampleWebServer/src/entities/user"
	"log"
)

func Create(newUser user.User) error {
	u := user.User{Name: "manager created", Password: "pass"}

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
	_, err = stmt.Exec(newUser.Name, newUser.Password)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
