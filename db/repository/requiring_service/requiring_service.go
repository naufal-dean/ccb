package required_service

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func Create(db *sql.DB, service, endpoint string) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("INSERT INTO requiring_service(service, endpoint) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(service, endpoint)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}

func GetServicesByEndpoint(db *sql.DB, endpoint string) (services []string) {
	stmt, err := db.Prepare("SELECT service FROM requiring_service WHERE endpoint = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var service string
		err = rows.Scan(&service)
		if err != nil {
			log.Fatal(err)
		}

		services = append(services, service)
	}
	return services
}
