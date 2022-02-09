package status

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func Create(db *sql.DB, service, endpoint, status string) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("INSERT INTO status(service, endpoint, status) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(service, endpoint, status)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}

func GetById(db *sql.DB, id int) (int, string, string, string) {
	stmt, err := db.Prepare("SELECT id, service, endpoint, status FROM status WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var service, endpoint, status string
	err = stmt.QueryRow(id).Scan(&id, &service, &endpoint, &status)
	if err != nil {
		log.Fatal(err)
	}
	return id, service, endpoint, status
}

func GetByServiceAndEndpoint(db *sql.DB, service, endpoint string) ([]int, []string, []string, []string) {
	stmt, err := db.Prepare("SELECT id, service, endpoint, status FROM status WHERE service = ? AND endpoint = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(service, endpoint)
	if err != nil {
		log.Fatal(err)
	}

	var ids []int
	var services, endpoints, statuses []string
	for rows.Next() {
		var id int
		var status string
		err = rows.Scan(&id, &service, &endpoint, &status)
		if err != nil {
			log.Fatal(err)
		}

		ids = append(ids, id)
		services = append(services, service)
		endpoints = append(endpoints, endpoint)
		statuses = append(statuses, status)
	}
	return ids, services, endpoints, statuses
}
