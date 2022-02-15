package server

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func Run(addr string, registerServer func(*grpc.Server)) {
	// Init listener
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

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

	grpcServer := grpc.NewServer(opts...)

	// Register HttpServer
	registerServer(grpcServer)

	// Serve
	log.Printf("Listening to: %s\n", addr)
	log.Fatal(grpcServer.Serve(lis))
}
