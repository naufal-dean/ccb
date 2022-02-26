package required_service

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/naufal-dean/ccb/lib/utils"
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

	stmt, err := tx.Prepare("INSERT OR IGNORE INTO required_service(endpoint, rd_service, rd_endpoint) VALUES(?, ?, ?)")
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
	stmt, err := sr.db.Prepare("SELECT DISTINCT endpoint FROM required_service WHERE rd_service = ?")
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

func (sr Sqlite3Repository) GetEndpointsByRdServiceAndRdEndpoints(rdServices, rdEndpoints []string) ([]string, error) {
	stmt, err := sr.db.Prepare("SELECT endpoint FROM required_service WHERE rd_service = ? AND rd_endpoint = ?")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()

	endpointsSet := make(map[string]bool)
	for i := range rdEndpoints {
		rows, err := stmt.Query(rdServices[i], rdEndpoints[i])
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

			endpointsSet[endpoint] = true
		}
	}
	return utils.GetMapKeys(endpointsSet), nil
}

func (sr Sqlite3Repository) oldGetEndpointsByRdServiceAndRdEndpoints(rdService string, rdEndpoints []string) ([]string, error) {
	// Deprecated
	stmt, err := sr.db.Prepare("SELECT endpoint FROM required_service WHERE rd_service = ? AND rd_endpoint = ?")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()

	endpointsSet := make(map[string]bool)
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

			endpointsSet[endpoint] = true
		}
	}
	return utils.GetMapKeys(endpointsSet), nil
}
