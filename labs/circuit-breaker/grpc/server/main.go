package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	// "google.golang.org/grpc/credentials"
	// "google.golang.org/grpc/examples/data"

	// "github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"

	pb "github.com/naufal-dean/ccb/labs/circuit-breaker/grpc/protobuf"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type usersServer struct {
	pb.UnimplementedUsersServer
}

func (s *usersServer) Register(ctx context.Context, user *pb.User) (*empty.Empty, error) {
	log.Println("Received request: Register")
	return new(empty.Empty), nil
}

func (s *usersServer) List(ctx context.Context, void *empty.Empty) (*pb.UserList, error) {
	log.Println("Received request: List")
	return &pb.UserList{
		List: []*pb.User{
			&pb.User{
				Name:  "test",
				Email: "test@test.com",
				Age:   20,
			},
			&pb.User{
				Name:  "test2",
				Email: "test2@test.com",
				Age:   202,
			},
		},
	}, nil
}

func (s *usersServer) Echo(ctx context.Context, data *wrappers.BytesValue) (*wrappers.BytesValue, error) {
	log.Println("Received request: Echo")
	return data, nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

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

	s := usersServer{}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterUsersServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
