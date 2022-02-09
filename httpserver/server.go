package httpserver

import (
	"github.com/naufal-dean/ccb/app"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net/http"

	pb "github.com/naufal-dean/ccb/protobuf"
)

type HttpServer struct {
	app app.App
	pb.UnimplementedHttpServer
}

func New(app app.App) *HttpServer {
	return &HttpServer{app: app}
}

func (s *HttpServer) Register(sr grpc.ServiceRegistrar) {
	pb.RegisterHttpServer(sr, s)
}

func (s *HttpServer) Request(ctx context.Context, input *pb.RequestInput) (*pb.Response, error) {
	log.Println("Received request: Request")

	res, err := doRequest(input.Method, input.Url, input.Body, input.Header)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return convertResponse(res)
}

func (s *HttpServer) Get(ctx context.Context, input *pb.GetInput) (*pb.Response, error) {
	log.Println("Received request: Get")

	res, err := doRequest(http.MethodGet, input.Url, nil, input.Header)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return convertResponse(res)
}

func (s *HttpServer) Post(ctx context.Context, input *pb.PostInput) (*pb.Response, error) {
	log.Println("Received request: Post")

	res, err := doRequest(http.MethodPost, input.Url, input.Body, input.Header)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return convertResponse(res)
}

func (s *HttpServer) Put(ctx context.Context, input *pb.PutInput) (*pb.Response, error) {
	log.Println("Received request: Put")

	res, err := doRequest(http.MethodPut, input.Url, input.Body, input.Header)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return convertResponse(res)
}

func (s *HttpServer) Delete(ctx context.Context, input *pb.DeleteInput) (*pb.Response, error) {
	log.Println("Received request: Delete")

	res, err := doRequest(http.MethodDelete, input.Url, nil, input.Header)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return convertResponse(res)
}
