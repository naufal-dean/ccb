package status

import (
	_ "github.com/mattn/go-sqlite3"
)

type Status struct {
	Service  string
	Endpoint string
	Status   string
}
