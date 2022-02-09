package required_service

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func Create(db *sql.DB, endpoint, service string) {
	// TODO: implement
}

func GetEndpointsByService(db *sql.DB, service string) (endpoints []string) {
	// TODO: implement
	return nil
}
