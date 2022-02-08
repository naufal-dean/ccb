package model

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	repository "github.com/naufal-dean/ccb/db/repository/status"
)

type Status struct {
	Id       int
	Service  string
	Endpoint string
	Status   string
}

func (s *Status) Create(db *sql.DB) {
	repository.Create(db, s.Service, s.Endpoint, s.Status)
}

func (s *Status) GetById(db *sql.DB, id int) *Status {
	service, endpoint, status := repository.GetById(db, id)
	s.Id = id
	s.Service = service
	s.Endpoint = endpoint
	s.Status = status
	return s
}
