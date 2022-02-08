package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/naufal-dean/ccb/db/repository"
	"github.com/naufal-dean/ccb/httpserver"
	"google.golang.org/grpc"
	"log"
	"net"
)

type App struct {
	Server *grpc.Server
	Db     *sql.DB
}

func initApp(a *App, port int, dbpath string) {
	initAppDb(a, dbpath)
	initAppServer(a)
}

func initAppDb(a *App, dbpath string) {
	var err error
	a.Db, err = sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Connection to db '%s' established...", dbpath)
}

func initAppServer(a *App) {
	// Create grpc server
	var opts []grpc.ServerOption

	// if *tls {
	// 	if *certFile == "" {
	// 		*certFile = data.Path("x509/server_cert.pem")
	// 	}
	// 	if *keyFile == "" {
	// 		*keyFile = data.Path("x509/server_key.pem")
	// 	}
	// 	creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
	// 	if err != nil {
	// 		log.Fatalf("Failed to generate credentials %v", err)
	// 	}
	// 	opts = []grpc.ServerOption{grpc.Creds(creds)}
	// }

	a.Server = grpc.NewServer(opts...)

	// Register HttpServer
	s := httpserver.HttpServer{}
	s.Register(a.Server)
}

func appExec(port int, dbpath string) {
	var a = &App{}

	// Init listener
	addr := fmt.Sprintf("localhost:%d", port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Init app
	initApp(a, port, dbpath)
	defer a.Db.Close()

	// TEST INSERT
	// TODO: remove
	repository.TestInsert(a.Db)

	// Serve
	log.Printf("Listening to: %s\n", addr)
	if err := a.Server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}

func main() {
	port := flag.Int("port", 50051, "The server port")
	dbpath := flag.String("dbpath", "", "The server sqlite3 db path")

	flag.Parse()

	appExec(*port, *dbpath)
}
