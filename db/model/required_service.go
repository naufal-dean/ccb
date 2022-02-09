package model

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	repository "github.com/naufal-dean/ccb/db/repository/required_service"
)

type RequiredService struct {
	Endpoint string
	Service  string
}

func (rs *RequiredService) Create(db *sql.DB) {
	repository.Create(db, rs.Endpoint, rs.Service)
}

func (rs *RequiredService) GetEndpointsByService(db *sql.DB, endpoint string) (endpoints []string) {
	return repository.GetEndpointsByService(db, endpoint)
}
