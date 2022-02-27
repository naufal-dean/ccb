package internal

import (
	"errors"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/naufal-dean/ccb/protobuf"
)

func BroadcastOpenCircuits(targetServiceAddr string, services, endpoints []string, expiry int64) error {
	log.Printf("Exec: BroadcastOpenCircuits: %s, %v, %v\n", targetServiceAddr, services, endpoints)

	if len(services) != len(endpoints) {
		return errors.New("services and endpoints array not in the same length")
	}

	conn, err := grpc.Dial(targetServiceAddr, grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to dial: %v\n", err)
		return errors.New("failed to dial")
	}
	defer conn.Close()

	client := pb.NewListenerClient(conn)

	_, err = client.OpenCircuits(context.Background(), &pb.ServiceEndpoints{
		Services:  services,
		Endpoints: endpoints,
		Expiry:    expiry,
	})
	if err != nil {
		log.Printf("Failed to request to server: %v\n", err)
		return errors.New("failed to request to server")
	}
	return nil
}

func BroadcastCloseCircuits(targetServiceAddr string, services, endpoints []string) error {
	log.Printf("Exec: BroadcastCloseCircuits: %s, %v, %v\n", targetServiceAddr, services, endpoints)

	if len(services) != len(endpoints) {
		return errors.New("services and endpoints array not in the same length")
	}

	conn, err := grpc.Dial(targetServiceAddr, grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to dial: %v\n", err)
		return errors.New("failed to dial")
	}
	defer conn.Close()

	client := pb.NewListenerClient(conn)

	_, err = client.CloseCircuits(context.Background(), &pb.ServiceEndpoints{
		Services:  services,
		Endpoints: endpoints,
	})
	if err != nil {
		log.Printf("Failed to request to server: %v\n", err)
		return errors.New("failed to request to server")
	}
	return nil
}
