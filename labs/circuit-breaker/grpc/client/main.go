package main

import (
	"bufio"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	// "google.golang.org/grpc/credentials"
	// "google.golang.org/grpc/examples/data"

	// "github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"

	pb "github.com/naufal-dean/labs/circuit-breaker/grpc/test"
)

var (
	serverAddr = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
)

func main() {
	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	c := pb.NewUsersClient(conn)

	userList, err := c.List(context.Background(), new(empty.Empty))
	if err != nil {
		log.Fatalf("Failed to request to server: %v", err)
	}

	data, _ := json.Marshal(userList)
	log.Printf("%s\n", data)

	  for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Request: ")
		text, _ := reader.ReadString('\n')

		echo, err := c.Echo(context.Background(), &wrappers.BytesValue{Value: []byte(text)})
		if err != nil {
			log.Fatalf("Failed to request to server: %v", err)
		}

		fmt.Printf("Response: %s", string(echo.Value))
	  }
}
