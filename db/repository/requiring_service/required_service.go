package required_service

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func Create(db *sql.DB, service, endpoint string) {
	// TODO: implement
}

func GetServicesByEndpoint(db *sql.DB, service string) (services []string) {
	// TODO: implement
	return nil
}
