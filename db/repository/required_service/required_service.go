package required_service

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func Create(db *sql.DB, endpoint, service string) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("INSERT INTO required_service(endpoint, service) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(endpoint, service)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}

func GetEndpointsByService(db *sql.DB, service string) (endpoints []string) {
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
