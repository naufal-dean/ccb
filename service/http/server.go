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

func (s HttpServer) storeDependency(targetUrl string, initialHeader map[string]string) (*string, *string, bool) {
	// Return parsed target endpoint to be reused
	u, err := url.Parse(targetUrl)
	if err != nil {
		log.Printf("Failed to parse target URL: %v\n", err)
		return nil, nil, false
	}

	rgServiceModel, rdServiceModel, ok := buildDependency(u.Host, u.Path, initialHeader)
	if !ok {
		log.Println("Skip storing dependency")
		return &u.Host, &u.Path, true
	}
	s.app.Repositories.RequiringService.Create(*rgServiceModel)
	s.app.Repositories.RequiredService.Create(*rdServiceModel)
	return &u.Host, &u.Path, true
}

func (s HttpServer) isCircuitOpen(method, service, endpoint string) bool {
	// Return true then API call will be cancelled, vice versa
	// TODO: add method column condition
	statusModels, err := s.app.Repositories.Status.GetByRdServiceAndRdEndpoint(service, endpoint)
	// If any db error, then allow the request (assume that the circuit closed)
	return err == nil && statusModels != nil
}

func (s HttpServer) Request(ctx context.Context, input *pb.RequestInput) (*pb.Response, error) {
	log.Println("Received request: Request")

	targetService, targetEndpoint, ok := s.storeDependency(input.Url, input.InitialHeader)
	if ok && s.isCircuitOpen(input.Method, *targetService, *targetEndpoint) {
		return new(pb.Response), status.Error(codes.Aborted, "Request cancelled due to opened circuit")
	}

	res, err := doRequest(input.Method, input.Url, input.Body, input.Header, &s.app.ServiceName, targetEndpoint)
	if err != nil {
		return new(pb.Response), status.Error(codes.Internal, "Failed to execute the request")
	}
	defer res.Body.Close()

	return convertResponse(res)
}

func (s HttpServer) Get(ctx context.Context, input *pb.GetInput) (*pb.Response, error) {
	log.Println("Received request: Get")

	targetService, targetEndpoint, ok := s.storeDependency(input.Url, input.InitialHeader)
	if ok && s.isCircuitOpen(http.MethodGet, *targetService, *targetEndpoint) {
		return new(pb.Response), status.Error(codes.Aborted, "Request cancelled due to opened circuit")
	}

	res, err := doRequest(http.MethodGet, input.Url, nil, input.Header, &s.app.ServiceName, targetEndpoint)
	if err != nil {
		return new(pb.Response), status.Error(codes.Internal, "Failed to execute the request")
	}
	defer res.Body.Close()

	return convertResponse(res)
}

func (s HttpServer) Post(ctx context.Context, input *pb.PostInput) (*pb.Response, error) {
	log.Println("Received request: Post")

	targetService, targetEndpoint, ok := s.storeDependency(input.Url, input.InitialHeader)
	if ok && s.isCircuitOpen(http.MethodPost, *targetService, *targetEndpoint) {
		return new(pb.Response), status.Error(codes.Aborted, "Request cancelled due to opened circuit")
	}

	res, err := doRequest(http.MethodPost, input.Url, input.Body, input.Header, &s.app.ServiceName, targetEndpoint)
	if err != nil {
		return new(pb.Response), status.Error(codes.Internal, "Failed to execute the request")
	}
	defer res.Body.Close()

	return convertResponse(res)
}

func (s HttpServer) Put(ctx context.Context, input *pb.PutInput) (*pb.Response, error) {
	log.Println("Received request: Put")

	targetService, targetEndpoint, ok := s.storeDependency(input.Url, input.InitialHeader)
	if ok && s.isCircuitOpen(http.MethodPut, *targetService, *targetEndpoint) {
		return new(pb.Response), status.Error(codes.Aborted, "Request cancelled due to opened circuit")
	}

	res, err := doRequest(http.MethodPut, input.Url, input.Body, input.Header, &s.app.ServiceName, targetEndpoint)
	if err != nil {
		return new(pb.Response), status.Error(codes.Internal, "Failed to execute the request")
	}
	defer res.Body.Close()

	return convertResponse(res)
}

func (s HttpServer) Delete(ctx context.Context, input *pb.DeleteInput) (*pb.Response, error) {
	log.Println("Received request: Delete")

	targetService, targetEndpoint, ok := s.storeDependency(input.Url, input.InitialHeader)
	if ok && s.isCircuitOpen(http.MethodDelete, *targetService, *targetEndpoint) {
		return new(pb.Response), status.Error(codes.Aborted, "Request cancelled due to opened circuit")
	}

	res, err := doRequest(http.MethodDelete, input.Url, nil, input.Header, &s.app.ServiceName, targetEndpoint)
	if err != nil {
		return new(pb.Response), status.Error(codes.Internal, "Failed to execute the request")
	}
	defer res.Body.Close()

	return convertResponse(res)
}
