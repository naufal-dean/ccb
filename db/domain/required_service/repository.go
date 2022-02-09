package required_service

import "database/sql"

type Repository interface {
	Create(db *sql.DB, model RequiredService)
	GetEndpointsByService(db *sql.DB, service string) (endpoints []string)
}
