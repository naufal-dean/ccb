package requiring_service

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

func (sr Sqlite3Repository) Create(model RequiringService) error {
	tx, err := sr.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}

	stmt, err := tx.Prepare("INSERT OR IGNORE INTO requiring_service(rg_service, endpoint) VALUES(?, ?)")
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(model.RgService, model.Endpoint)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (sr Sqlite3Repository) GetRgServicesByEndpoint(endpoint string) ([]string, error) {
	stmt, err := sr.db.Prepare("SELECT rg_service FROM requiring_service WHERE endpoint = ?")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(endpoint)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var rgServices []string
	for rows.Next() {
		var rgService string
		err = rows.Scan(&rgService)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		rgServices = append(rgServices, rgService)
	}
	return rgServices, nil
}

func (sr Sqlite3Repository) GetDependencyMapByEndpoints(endpoints []string) (map[string][]string, error) {
	stmt, err := sr.db.Prepare("SELECT rg_service FROM requiring_service WHERE endpoint = ?")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()

	dependencyMap := make(map[string][]string)
	for _, endpoint := range endpoints {
		rows, err := stmt.Query(endpoint)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		for rows.Next() {
			var rgService string
			err = rows.Scan(&rgService)
			if err != nil {
				log.Println(err)
				return nil, err
			}

			if _, ok := dependencyMap[rgService]; ok {
				dependencyMap[rgService] = append(dependencyMap[rgService], rgService)
			} else {
				dependencyMap[rgService] = []string{endpoint}
			}
		}
	}
	return dependencyMap, nil
}
