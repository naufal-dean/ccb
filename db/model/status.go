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
	id, service, endpoint, status := repository.GetById(db, id)
	s.Id = id
	s.Service = service
	s.Endpoint = endpoint
	s.Status = status
	return s
}

func (s *Status) GetByServiceAndEndpoint(db *sql.DB, service, endpoint string) []*Status {
	ids, services, endpoints, statuses := repository.GetByServiceAndEndpoint(db, service, endpoint)
	var res []*Status
	for i := range ids {
		res = append(res, &Status{
			Id:       ids[i],
			Service:  services[i],
			Endpoint: endpoints[i],
			Status:   statuses[i],
		})
	}
	return res
}
