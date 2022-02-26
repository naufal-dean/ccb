package internal

import (
	"errors"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/naufal-dean/ccb/protobuf"
)

func BroadcastOpenCircuits(currServiceName, targetServiceAddr string, endpoints []string) error {
	log.Printf("Exec: BroadcastOpenCircuits: %s, %s, %v\n", currServiceName, targetServiceAddr, endpoints)

	conn, err := grpc.Dial(targetServiceAddr, grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to dial: %v\n", err)
		return errors.New("failed to dial")
	}
	defer conn.Close()

	client := pb.NewListenerClient(conn)

	_, err = client.OpenCircuits(context.Background(), &pb.ServiceEndpoints{
		Service:   currServiceName,
		Endpoints: endpoints,
	})
	if err != nil {
		log.Printf("Failed to request to server: %v\n", err)
		return errors.New("failed to request to server")
	}
	return nil
}

func BroadcastCloseCircuits(currServiceName, targetServiceAddr string, endpoints []string) error {
	log.Printf("Exec: BroadcastCloseCircuits: %s, %s, %v\n", currServiceName, targetServiceAddr, endpoints)

	conn, err := grpc.Dial(targetServiceAddr, grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to dial: %v\n", err)
		return errors.New("failed to dial")
	}
	defer conn.Close()

	client := pb.NewListenerClient(conn)

	_, err = client.CloseCircuits(context.Background(), &pb.ServiceEndpoints{
		Service:   currServiceName,
		Endpoints: endpoints,
	})
	if err != nil {
		log.Printf("Failed to request to server: %v\n", err)
		return errors.New("failed to request to server")
	}
	return nil
}
