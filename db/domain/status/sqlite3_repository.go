package status

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

func (sr Sqlite3Repository) Create(model Status) {
	tx, err := sr.db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("INSERT INTO status(rd_service, rd_endpoint, status) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(model.RdService, model.RdEndpoint, model.Status)
	if err != nil {
		log.Fatal(err)
	}

	tx.Commit()
}

func (sr Sqlite3Repository) CreateFromOneRdServiceAndManyRdEndpoints(rdService string, rdEndpoints []string, status string) {
	tx, err := sr.db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// Use insert or ignore to assert uniqueness
	stmt, err := tx.Prepare("INSERT OR IGNORE INTO status(rd_service, rd_endpoint, status) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, rdEndpoint := range rdEndpoints {
		_, err = stmt.Exec(rdService, rdEndpoint, status)
		if err != nil {
			log.Fatal(err)
		}
	}

	tx.Commit()
}

func (sr Sqlite3Repository) GetByRdServiceAndRdEndpoint(rdService, rdEndpoint string) []*Status {
	stmt, err := sr.db.Prepare("SELECT rd_service, rd_endpoint, status FROM status WHERE rd_service = ? AND rd_endpoint = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(rdService, rdEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	var models []*Status
	for rows.Next() {
		model := &Status{}
		err = rows.Scan(&model.RdService, &model.RdEndpoint, &model.Status)
		if err != nil {
			log.Fatal(err)
		}

		models = append(models, model)
	}
	return models
}

func (sr Sqlite3Repository) DeleteWhereOneRdServiceAndManyRdEndpointsEqual(rdService string, rdEndpoints []string) {
	tx, err := sr.db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("DELETE FROM status WHERE rd_service = ? AND rd_endpoint = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, rdEndpoint := range rdEndpoints {
		_, err = stmt.Exec(rdService, rdEndpoint)
		if err != nil {
			log.Fatal(err)
		}
	}

	tx.Commit()
}
