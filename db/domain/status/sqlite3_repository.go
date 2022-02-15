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

func (sr Sqlite3Repository) Create(model Status) error {
	tx, err := sr.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO status(rd_service, rd_endpoint, status) VALUES(?, ?, ?)")
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(model.RdService, model.RdEndpoint, model.Status)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (sr Sqlite3Repository) CreateFromOneRdServiceAndManyRdEndpoints(rdService string, rdEndpoints []string, status string) ([]string, error) {
	tx, err := sr.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	checkStmt, err := tx.Prepare("SELECT 1 FROM status WHERE rd_service = ? AND rd_endpoint = ?")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer checkStmt.Close()

	stmt, err := tx.Prepare("INSERT INTO status(rd_service, rd_endpoint, status) VALUES(?, ?, ?)")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()

	var createdRdEndpoints []string
	for _, rdEndpoint := range rdEndpoints {
		// If row already exists, then there are no status update (no need to broadcast later)
		// If the query selects no rows, the *Row's Scan will return ErrNoRows
		var temp int
		err := checkStmt.QueryRow(rdService, rdEndpoint).Scan(&temp)
		if err == nil && temp == 1 {
			log.Printf("Endpoint %s not updated\n", rdEndpoint) // Debug
			continue
		}
		// Here the row to be created is guaranteed to be unique (PK pair not already exists)
		_, err = stmt.Exec(rdService, rdEndpoint, status)
		if err != nil {
			log.Println(err)
			tx.Rollback()
			return nil, err
		}
		createdRdEndpoints = append(createdRdEndpoints, rdEndpoint)
	}

	tx.Commit()
	return createdRdEndpoints, nil
}

func (sr Sqlite3Repository) GetByRdServiceAndRdEndpoint(rdService, rdEndpoint string) ([]*Status, error) {
	stmt, err := sr.db.Prepare("SELECT rd_service, rd_endpoint, status FROM status WHERE rd_service = ? AND rd_endpoint = ?")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(rdService, rdEndpoint)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var models []*Status
	for rows.Next() {
		model := &Status{}
		err = rows.Scan(&model.RdService, &model.RdEndpoint, &model.Status)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		models = append(models, model)
	}
	return models, nil
}

func (sr Sqlite3Repository) DeleteWhereOneRdServiceAndManyRdEndpointsEqual(rdService string, rdEndpoints []string) ([]string, error) {
	tx, err := sr.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	checkStmt, err := tx.Prepare("SELECT 1 FROM status WHERE rd_service = ? AND rd_endpoint = ?")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer checkStmt.Close()

	stmt, err := tx.Prepare("DELETE FROM status WHERE rd_service = ? AND rd_endpoint = ?")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()

	var deletedRdEndpoints []string
	for _, rdEndpoint := range rdEndpoints {
		// If row not exists, then there are no status update (no need to broadcast later)
		// If the query selects no rows, the *Row's Scan will return ErrNoRows
		var temp int
		err := checkStmt.QueryRow(rdService, rdEndpoint).Scan(&temp)
		if err != nil {
			// TODO: maybe check the error type, assert that it is ErrNoRows
			log.Printf("Endpoint %s not updated\n", rdEndpoint) // Debug
			continue
		}
		// Here the row to be deleted is guaranteed to be exists
		_, err = stmt.Exec(rdService, rdEndpoint)
		if err != nil {
			log.Println(err)
			tx.Rollback()
			return nil, err
		}
		deletedRdEndpoints = append(deletedRdEndpoints, rdEndpoint)
	}

	tx.Commit()
	return deletedRdEndpoints, nil
}

func (sr Sqlite3Repository) oldCreateFromOneRdServiceAndManyRdEndpoints(rdService string, rdEndpoints []string, status string) error {
	// Deprecated
	tx, err := sr.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}

	// Use insert or ignore to assert uniqueness
	stmt, err := tx.Prepare("INSERT OR IGNORE INTO status(rd_service, rd_endpoint, status) VALUES(?, ?, ?)")
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmt.Close()

	for _, rdEndpoint := range rdEndpoints {
		_, err = stmt.Exec(rdService, rdEndpoint, status)
		if err != nil {
			log.Println(err)
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

func (sr Sqlite3Repository) oldDeleteWhereOneRdServiceAndManyRdEndpointsEqual(rdService string, rdEndpoints []string) error {
	// Deprecated
	tx, err := sr.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}

	stmt, err := tx.Prepare("DELETE FROM status WHERE rd_service = ? AND rd_endpoint = ?")
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmt.Close()

	for _, rdEndpoint := range rdEndpoints {
		_, err = stmt.Exec(rdService, rdEndpoint)
		if err != nil {
			log.Println(err)
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}
