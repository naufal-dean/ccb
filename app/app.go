package app

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"log"
)

type App struct {
	Server *grpc.Server
	Db     *sql.DB
}

func initAppDb(a *App, dbpath string) {
	var err error
	a.Db, err = sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Connection to db '%s' established...", dbpath)
}

func New(dbpath string) App {
	application := App{}

	initAppDb(&application, dbpath)

	return application
}
