package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"google.golang.org/grpc"

	"github.com/naufal-dean/ccb/app"
	"github.com/naufal-dean/ccb/server"
	httpservice "github.com/naufal-dean/ccb/service/http"
	listenerservice "github.com/naufal-dean/ccb/service/listener"
)

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value, ok := os.LookupEnv(key); ok {
		valueInt, err := strconv.Atoi(value)
		if err != nil {
			return defaultValue
		}
		return valueInt
	}
	return defaultValue
}

func main() {
	// TODO: load env file config if needed
	port := flag.Int("port", getEnvInt("CCB_PORT", 50051), "The server port")
	serviceName := flag.String("name", getEnv("CCB_SERVICE_NAME", "localhost:50051"), "This service name on Kube DNS")
	dbPath := flag.String("dbpath", getEnv("CCB_DB_PATH", ""), "The server sqlite3 db path")

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
