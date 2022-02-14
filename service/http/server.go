package http

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

func (s HttpServer) storeDependency(targetUrl string, initialHeader map[string]string) {
	rgServiceModel, rdServiceModel, ok := buildDependency(targetUrl, initialHeader)
	if !ok {
		log.Println("Skip storing dependency")
		return
	}
	s.app.Repositories.RequiringService.Create(*rgServiceModel)
	s.app.Repositories.RequiredService.Create(*rdServiceModel)
}

func (s HttpServer) Request(ctx context.Context, input *pb.RequestInput) (*pb.Response, error) {
	log.Println("Received request: Request")

	s.storeDependency(input.Url, input.InitialHeader)

	res, err := doRequest(input.Method, input.Url, input.Body, input.Header)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return convertResponse(res)
}

func (s HttpServer) Get(ctx context.Context, input *pb.GetInput) (*pb.Response, error) {
	log.Println("Received request: Get")

	s.storeDependency(input.Url, input.InitialHeader)

	res, err := doRequest(http.MethodGet, input.Url, nil, input.Header)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return convertResponse(res)
}

func (s HttpServer) Post(ctx context.Context, input *pb.PostInput) (*pb.Response, error) {
	log.Println("Received request: Post")

	s.storeDependency(input.Url, input.InitialHeader)

	res, err := doRequest(http.MethodPost, input.Url, input.Body, input.Header)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return convertResponse(res)
}

func (s HttpServer) Put(ctx context.Context, input *pb.PutInput) (*pb.Response, error) {
	log.Println("Received request: Put")

	s.storeDependency(input.Url, input.InitialHeader)

	res, err := doRequest(http.MethodPut, input.Url, input.Body, input.Header)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return convertResponse(res)
}

func (s HttpServer) Delete(ctx context.Context, input *pb.DeleteInput) (*pb.Response, error) {
	log.Println("Received request: Delete")

	s.storeDependency(input.Url, input.InitialHeader)

	res, err := doRequest(http.MethodDelete, input.Url, nil, input.Header)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return convertResponse(res)
}
