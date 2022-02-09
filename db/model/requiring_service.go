package model

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	repository "github.com/naufal-dean/ccb/db/repository/requiring_service"
)

type RequiringService struct {
	Service  string
	Endpoint string
}

func (rs *RequiringService) Create(db *sql.DB) {
	repository.Create(db, rs.Service, rs.Endpoint)
}

func (rs *RequiringService) GetServicesByEndpoint(db *sql.DB, endpoint string) (services []string) {
	return repository.GetServicesByEndpoint(db, endpoint)
}
