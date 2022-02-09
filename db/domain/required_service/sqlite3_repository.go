package required_service

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Sqlite3Repository struct {
	db *sql.DB
}

func NewSqlite3Repository(db *sql.DB) Sqlite3Repository {
	return Sqlite3Repository{db: db}
}

func (sr Sqlite3Repository) Create(db *sql.DB, model RequiredService) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("INSERT INTO required_service(endpoint, service) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(model.Endpoint, model.Service)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}

func (sr Sqlite3Repository) GetEndpointsByService(db *sql.DB, service string) (endpoints []string) {
	stmt, err := db.Prepare("SELECT endpoint FROM required_service WHERE service = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(service)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var endpoint string
		err = rows.Scan(&endpoint)
		if err != nil {
			log.Fatal(err)
		}

		endpoints = append(endpoints, endpoint)
	}
	return endpoints
}
