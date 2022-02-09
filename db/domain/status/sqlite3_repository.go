package status

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

func (sr Sqlite3Repository) Create(model Status) {
	tx, err := sr.db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("INSERT INTO status(service, endpoint, status) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(model.Service, model.Endpoint, model.Status)
	if err != nil {
		log.Fatal(err)
	}

	tx.Commit()
}

func (sr Sqlite3Repository) CreateFromOneServiceAndManyEndpoints(service string, endpoints []string, status string) {
	tx, err := sr.db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// Use insert or ignore to assert uniqueness
	stmt, err := tx.Prepare("INSERT OR IGNORE INTO status(service, endpoint, status) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, endpoint := range endpoints {
		_, err = stmt.Exec(service, endpoint, status)
		if err != nil {
			log.Fatal(err)
		}
	}

	tx.Commit()
}

func (sr Sqlite3Repository) GetById(id int) *Status {
	stmt, err := sr.db.Prepare("SELECT id, service, endpoint, status FROM status WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var service, endpoint, status string
	err = stmt.QueryRow(id).Scan(&id, &service, &endpoint, &status)
	if err != nil {
		log.Fatal(err)
	}
	return &Status{
		Id:       id,
		Service:  service,
		Endpoint: endpoint,
		Status:   status,
	}
}

func (sr Sqlite3Repository) GetByServiceAndEndpoint(service, endpoint string) []*Status {
	stmt, err := sr.db.Prepare("SELECT id, service, endpoint, status FROM status WHERE service = ? AND endpoint = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(service, endpoint)
	if err != nil {
		log.Fatal(err)
	}

	var models []*Status
	for rows.Next() {
		var id int
		var status string
		err = rows.Scan(&id, &service, &endpoint, &status)
		if err != nil {
			log.Fatal(err)
		}

		models = append(models, &Status{
			Id:       id,
			Service:  service,
			Endpoint: endpoint,
			Status:   status,
		})
	}
	return models
}
