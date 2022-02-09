package main

import (
	"flag"
	pb "github.com/naufal-dean/ccb/examples/client/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

var (
	serverAddr = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
)

func testGet(c pb.HttpClient) {
	response, err := c.Get(context.Background(), &pb.GetInput{
		Url:    "https://api.sampleapis.com/coffee/hot",
		Header: nil,
	})
	if err != nil {
		log.Fatalf("Failed to request to server: %v", err)
	}
	log.Println(string(response.Body))
}

func testOpenCircuits(c pb.ListenerClient) {
	response, err := c.OpenCircuits(context.Background(), &pb.ServiceEndpoints{
		Service: "test-euy",
		Endpoints: []string{
			"/124",
			"/hehe",
			"/123",
			"/random-euy",
		},
	})
	if err != nil {
		log.Fatalf("Failed to request to server: %v", err)
	}
	log.Println(response)
}

func testCloseCircuits(c pb.ListenerClient) {
	response, err := c.CloseCircuits(context.Background(), &pb.ServiceEndpoints{
		Service: "test-euy",
		Endpoints: []string{
			"/123",
			"/random-euy",
		},
	})
	if err != nil {
		log.Fatalf("Failed to request to server: %v", err)
	}
	log.Println(response)
}

func main() {
	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	hc := pb.NewHttpClient(conn)
	testGet(hc)

	lc := pb.NewListenerClient(conn)
	testOpenCircuits(lc)
	testCloseCircuits(lc)
}
