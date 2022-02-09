package main

import (
	"flag"
	"fmt"
	"github.com/naufal-dean/ccb/app"
	"github.com/naufal-dean/ccb/httpserver"
	"github.com/naufal-dean/ccb/listenerserver"
	"github.com/naufal-dean/ccb/server"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 50051, "The server port")
	dbpath := flag.String("dbpath", "", "The server sqlite3 db path")

	flag.Parse()
	addr := fmt.Sprintf("localhost:%d", *port)

	application := app.New(*dbpath)

	server.Run(addr, func(grpcServer *grpc.Server) {
		// Register HttpServer
		httpServer := httpserver.New(application)
		httpServer.Register(grpcServer)
		// Register ListenerServer
		listenerServer := listenerserver.New(application)
		listenerServer.Register(grpcServer)
	})
}
