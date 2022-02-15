package http

import (
	"log"
	"net/http"
	"net/url"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/naufal-dean/ccb/app"
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

func (s HttpServer) storeDependency(targetUrl string, initialHeader map[string]string) *string {
	// Return parsed target endpoint to be reused
	u, err := url.Parse(targetUrl)
	if err != nil {
		log.Printf("Failed to parse target URL: %v\n", err)
		return nil
	}

	rgServiceModel, rdServiceModel, ok := buildDependency(u.Host, u.Path, initialHeader)
	if !ok {
		log.Println("Skip storing dependency")
		return &u.Path
	}
	s.app.Repositories.RequiringService.Create(*rgServiceModel)
	s.app.Repositories.RequiredService.Create(*rdServiceModel)
	return &u.Path
}

func (s HttpServer) Request(ctx context.Context, input *pb.RequestInput) (*pb.Response, error) {
	log.Println("Received request: Request")

	targetEndpoint := s.storeDependency(input.Url, input.InitialHeader)

	res, err := doRequest(input.Method, input.Url, input.Body, input.Header, &s.app.ServiceName, targetEndpoint)
	if err != nil {
		return new(pb.Response), status.Error(codes.Internal, "Failed to execute the request")
	}
	defer res.Body.Close()

	return convertResponse(res)
}

func (s HttpServer) Get(ctx context.Context, input *pb.GetInput) (*pb.Response, error) {
	log.Println("Received request: Get")

	targetEndpoint := s.storeDependency(input.Url, input.InitialHeader)

	res, err := doRequest(http.MethodGet, input.Url, nil, input.Header, &s.app.ServiceName, targetEndpoint)
	if err != nil {
		return new(pb.Response), status.Error(codes.Internal, "Failed to execute the request")
	}
	defer res.Body.Close()

	return convertResponse(res)
}

func (s HttpServer) Post(ctx context.Context, input *pb.PostInput) (*pb.Response, error) {
	log.Println("Received request: Post")

	targetEndpoint := s.storeDependency(input.Url, input.InitialHeader)

	res, err := doRequest(http.MethodPost, input.Url, input.Body, input.Header, &s.app.ServiceName, targetEndpoint)
	if err != nil {
		return new(pb.Response), status.Error(codes.Internal, "Failed to execute the request")
	}
	defer res.Body.Close()

	return convertResponse(res)
}

func (s HttpServer) Put(ctx context.Context, input *pb.PutInput) (*pb.Response, error) {
	log.Println("Received request: Put")

	targetEndpoint := s.storeDependency(input.Url, input.InitialHeader)

	res, err := doRequest(http.MethodPut, input.Url, input.Body, input.Header, &s.app.ServiceName, targetEndpoint)
	if err != nil {
		return new(pb.Response), status.Error(codes.Internal, "Failed to execute the request")
	}
	defer res.Body.Close()

	return convertResponse(res)
}

func (s HttpServer) Delete(ctx context.Context, input *pb.DeleteInput) (*pb.Response, error) {
	log.Println("Received request: Delete")

	targetEndpoint := s.storeDependency(input.Url, input.InitialHeader)

	res, err := doRequest(http.MethodDelete, input.Url, nil, input.Header, &s.app.ServiceName, targetEndpoint)
	if err != nil {
		return new(pb.Response), status.Error(codes.Internal, "Failed to execute the request")
	}
	defer res.Body.Close()

	return convertResponse(res)
}
