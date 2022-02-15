package internal

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/naufal-dean/ccb/protobuf"
)

func BroadcastOpenCircuits(currServiceName, targetServiceAddr string, endpoints []string) {
	log.Println("Exec: BroadcastOpenCircuits")

	conn, err := grpc.Dial(targetServiceAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewListenerClient(conn)

	_, err = client.OpenCircuits(context.Background(), &pb.ServiceEndpoints{
		Service:   currServiceName,
		Endpoints: endpoints,
	})
	if err != nil {
		log.Fatalf("Failed to request to server: %v", err)
	}
}

func BroadcastCloseCircuits(currServiceName, targetServiceAddr string, endpoints []string) {
	log.Println("Exec: BroadcastCloseCircuits")

	conn, err := grpc.Dial(targetServiceAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewListenerClient(conn)

	_, err = client.CloseCircuits(context.Background(), &pb.ServiceEndpoints{
		Service:   currServiceName,
		Endpoints: endpoints,
	})
	if err != nil {
		log.Fatalf("Failed to request to server: %v", err)
	}
}
