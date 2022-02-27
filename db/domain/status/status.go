package status

import (
	_ "github.com/mattn/go-sqlite3"
)

type Status struct {
	RdService  string
	RdEndpoint string
	Status     string
	Expiry     int64
}
