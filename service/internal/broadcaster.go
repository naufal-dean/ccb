package internal

import (
	pb "github.com/naufal-dean/ccb/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

const (
	// TODO: get service name from env var or app struct
	serviceName = "service-test"
)

func BroadcastOpenCircuits(serverAddr string, endpoints []string) {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewListenerClient(conn)

	_, err = client.OpenCircuits(context.Background(), &pb.ServiceEndpoints{
		Service:   serviceName,
		Endpoints: endpoints,
	})
	if err != nil {
		log.Fatalf("Failed to request to server: %v", err)
	}
}

func BroadcastCloseCircuits(serverAddr string, endpoints []string) {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewListenerClient(conn)

	_, err = client.CloseCircuits(context.Background(), &pb.ServiceEndpoints{
		Service:   serviceName,
		Endpoints: endpoints,
	})
	if err != nil {
		log.Fatalf("Failed to request to server: %v", err)
	}
}
