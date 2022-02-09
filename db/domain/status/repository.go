package status

import "database/sql"

type Repository interface {
	Create(db *sql.DB, model Status)
	GetById(db *sql.DB, id int) *Status
	GetByServiceAndEndpoint(db *sql.DB, service, endpoint string) []*Status
}
