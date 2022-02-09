package requiring_service

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

func (sr Sqlite3Repository) Create(model RequiringService) {
	tx, err := sr.db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("INSERT INTO requiring_service(service, endpoint) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(model.Service, model.Endpoint)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}

func (sr Sqlite3Repository) GetServicesByEndpoint(endpoint string) (services []string) {
	stmt, err := sr.db.Prepare("SELECT service FROM requiring_service WHERE endpoint = ?")
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
