package main

import (
	"encoding/json"
	"flag"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"

	pb "github.com/naufal-dean/ccb/examples/client/protobuf"
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

	c := pb.NewHttpClient(conn)

	userList, err := c.Get(context.Background(), &pb.GetInput{})
	if err != nil {
		log.Fatalf("Failed to request to server: %v", err)
	}

	data, _ := json.Marshal(userList)
	log.Printf("%s\n", data)

	//for {
	//	reader := bufio.NewReader(os.Stdin)
	//	fmt.Print("Request: ")
	//	text, _ := reader.ReadString('\n')
	//
	//	echo, err := c.Echo(context.Background(), &wrappers.BytesValue{Value: []byte(text)})
	//	if err != nil {
	//		log.Fatalf("Failed to request to server: %v", err)
	//	}
	//
	//	fmt.Printf("Response: %s", string(echo.Value))
	//}
}
