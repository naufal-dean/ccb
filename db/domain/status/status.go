package status

import (
	_ "github.com/mattn/go-sqlite3"
)

type Status struct {
	Id       int
	Service  string
	Endpoint string
	Status   string
}
