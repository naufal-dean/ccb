package httpserver

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net/http"

	pb "github.com/naufal-dean/ccb/protobuf"
)

type HttpServer struct {
	pb.UnimplementedHttpServer
}

func (s *HttpServer) Register(sr grpc.ServiceRegistrar) {
	pb.RegisterHttpServer(sr, s)
}

func (s *HttpServer) Request(ctx context.Context, input *pb.RequestInput) (*pb.Response, error) {
	log.Println("Received request: Request")

	return &pb.Response{}, nil
}

func (s *HttpServer) Get(ctx context.Context, input *pb.GetInput) (*pb.Response, error) {
	log.Println("Received request: Get")

	response, err := http.Get(input.Url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := readBody(response)
	if err != nil {
		return nil, err
	}

	return convertResponse(response, body), nil
}

func (s *HttpServer) Post(ctx context.Context, input *pb.PostInput) (*pb.Response, error) {
	log.Println("Received request: Post")

	return &pb.Response{}, nil
}
