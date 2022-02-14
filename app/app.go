package app

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/naufal-dean/ccb/db/domain/required_service"
	"github.com/naufal-dean/ccb/db/domain/requiring_service"
	"github.com/naufal-dean/ccb/db/domain/status"
	"log"
)

type App struct {
	ServiceName  string
	Repositories Repositories
}

type Repositories struct {
	RequiredService  required_service.Repository
	RequiringService requiring_service.Repository
	Status           status.Repository
}

func initAppDb(dbPath string) *sql.DB {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Connection to db '%s' established...", dbPath)
	return db
}

func New(serviceName, dbPath string) App {
	db := initAppDb(dbPath)

	return App{
		ServiceName: serviceName,
		Repositories: Repositories{
			RequiredService:  required_service.NewSqlite3Repository(db),
			RequiringService: requiring_service.NewSqlite3Repository(db),
			Status:           status.NewSqlite3Repository(db),
		},
	}
}
