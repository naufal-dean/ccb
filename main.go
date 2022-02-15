package main

import (
	"flag"
	"fmt"

	"google.golang.org/grpc"

	"github.com/naufal-dean/ccb/app"
	"github.com/naufal-dean/ccb/server"
	httpservice "github.com/naufal-dean/ccb/service/http"
	listenerservice "github.com/naufal-dean/ccb/service/listener"
)

func main() {
	port := flag.Int("port", 50051, "The server port")
	// TODO: get service name from env config
	serviceName := flag.String("name", "localhost:50051", "This service name on Kube DNS")
	dbPath := flag.String("dbpath", "", "The server sqlite3 db path")

	flag.Parse()
	addr := fmt.Sprintf(":%d", *port)

	application := app.New(*serviceName, *dbPath)

	// TEST
	// TODO: remove
	//application.Repositories.Status.Create(status.Status{
	//	Service:  "test-euy",
	//	Endpoint: "/123",
	//	Status:   "CLOSED",
	//})
	//depMap := application.Repositories.RequiringService.GetDependencyMapByEndpoints([]string{"/endpoint-aa", "/endpoint-d"})
	//log.Print(depMap)
	// TEST END

	server.Run(addr, func(grpcServer *grpc.Server) {
		// Register service Http
		httpService := httpservice.New(application)
		httpService.Register(grpcServer)
		// Register service Listener
		listenerService := listenerservice.New(application)
		listenerService.Register(grpcServer)
	})
}
