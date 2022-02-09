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
	Repositories Repositories
}

type Repositories struct {
	requiredServiceRepository  required_service.Repository
	requiringServiceRepository requiring_service.Repository
	statusRepository           status.Repository
}

func initAppDb(dbpath string) *sql.DB {
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Connection to db '%s' established...", dbpath)
	return db
}

func New(dbpath string) App {
	db := initAppDb(dbpath)

	return App{
		Repositories: Repositories{
			requiredServiceRepository:  required_service.NewSqlite3Repository(db),
			requiringServiceRepository: requiring_service.NewSqlite3Repository(db),
			statusRepository:           status.NewSqlite3Repository(db),
		},
	}
}
