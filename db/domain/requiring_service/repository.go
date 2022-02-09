package requiring_service

import "database/sql"

type Repository interface {
	Create(db *sql.DB, model RequiringService)
	GetServicesByEndpoint(db *sql.DB, endpoint string) (services []string)
}
