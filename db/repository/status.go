package repository

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func TestInsert(db *sql.DB) {
	// TODO: remove

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("INSERT INTO status(service, endpoint, status) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec("service-a", "/coffee/make", "OPEN")
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}
