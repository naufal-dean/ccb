package required_service

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Sqlite3Repository struct {
	db *sql.DB
}

func NewSqlite3Repository(db *sql.DB) Sqlite3Repository {
	return Sqlite3Repository{db: db}
}

func (sr Sqlite3Repository) Create(model RequiredService) error {
	tx, err := sr.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO required_service(endpoint, rd_service, rd_endpoint) VALUES(?, ?, ?)")
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(model.Endpoint, model.RdService, model.RdEndpoint)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (sr Sqlite3Repository) GetEndpointsByRdService(rdService string) ([]string, error) {
	stmt, err := sr.db.Prepare("SELECT endpoint FROM required_service WHERE rd_service = ?")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(rdService)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var endpoints []string
	for rows.Next() {
		var endpoint string
		err = rows.Scan(&endpoint)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		endpoints = append(endpoints, endpoint)
	}
	return endpoints, nil
}

func (sr Sqlite3Repository) GetEndpointsByRdServiceAndRdEndpoints(rdService string, rdEndpoints []string) ([]string, error) {
	stmt, err := sr.db.Prepare("SELECT endpoint FROM required_service WHERE rd_service = ? AND rd_endpoint = ?")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()

	var endpoints []string
	for _, rdEndpoint := range rdEndpoints {
		rows, err := stmt.Query(rdService, rdEndpoint)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		for rows.Next() {
			var endpoint string
			err = rows.Scan(&endpoint)
			if err != nil {
				log.Println(err)
				return nil, err
			}

			endpoints = append(endpoints, endpoint)
		}
	}
	return endpoints, nil
}
