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

	stmt, err := tx.Prepare("INSERT INTO requiring_service(rg_service, endpoint) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(model.RgService, model.Endpoint)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}

func (sr Sqlite3Repository) GetRgServicesByEndpoint(endpoint string) (rgServices []string) {
	stmt, err := sr.db.Prepare("SELECT rg_service FROM requiring_service WHERE endpoint = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var rgService string
		err = rows.Scan(&rgService)
		if err != nil {
			log.Fatal(err)
		}

		rgServices = append(rgServices, rgService)
	}
	return rgServices
}

func (sr Sqlite3Repository) GetDependencyMapByEndpoints(endpoints []string) map[string][]string {
	stmt, err := sr.db.Prepare("SELECT rg_service FROM requiring_service WHERE endpoint = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	dependencyMap := make(map[string][]string)
	for _, endpoint := range endpoints {
		rows, err := stmt.Query(endpoint)
		if err != nil {
			log.Fatal(err)
		}

		for rows.Next() {
			var rgService string
			err = rows.Scan(&rgService)
			if err != nil {
				log.Fatal(err)
			}

			if _, ok := dependencyMap[rgService]; ok {
				dependencyMap[rgService] = append(dependencyMap[rgService], rgService)
			} else {
				dependencyMap[rgService] = []string{endpoint}
			}
		}
	}
	return dependencyMap
}
