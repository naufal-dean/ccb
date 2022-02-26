package http

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/naufal-dean/ccb/app"
	"github.com/naufal-dean/ccb/lib/circuitbreaker"
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

func (s HttpServer) storeDependency(targetUrl string, initialHeader map[string]string) (*url.URL, bool) {
	// Return parsed target endpoint to be reused
	u, err := url.Parse(targetUrl)
	if err != nil {
		log.Printf("Failed to parse target URL: %v\n", err)
		return nil, false
	}

	rgServiceModel, rdServiceModel, ok := buildDependency(u.Host, u.Path, initialHeader)
	if !ok {
		log.Println("Skip storing dependency")
		return u, true
	}
	s.app.Repositories.RequiringService.Create(*rgServiceModel)
	s.app.Repositories.RequiredService.Create(*rdServiceModel)
	return u, true
}

func (s HttpServer) isCircuitOpen(method, service, endpoint string) bool {
	// Return true then API call will be cancelled, vice versa
	statusModels, err := s.app.Repositories.Status.GetByRdServiceAndRdEndpoint(service, fmt.Sprintf("%s:%s", method, endpoint))
	// If any db error, then allow the request (assume that the circuit closed)
	return err == nil && statusModels != nil
}

func (s HttpServer) Request(ctx context.Context, input *pb.RequestInput) (*pb.Response, error) {
	log.Println("Received request: Request")

	reqUrl, ok := s.storeDependency(input.Url, input.InitialHeader)
	if ok && s.isCircuitOpen(input.Method, reqUrl.Host, reqUrl.Path) {
		return new(pb.Response), status.Error(codes.Aborted, "Request cancelled due to opened circuit")
	}

	res, err := s.doRequest(input.Method, input.Url, input.Body, input.Header, &s.app.ServiceName, reqUrl)
	if err != nil {
		if errors.Is(err, circuitbreaker.ErrOpenState) {
			return new(pb.Response), status.Error(codes.Aborted, "Request cancelled due to opened circuit")
		}
		return new(pb.Response), status.Error(codes.Internal, "Failed to execute the request")
	}
	defer res.Body.Close()

	return convertResponse(res)
}

func (s HttpServer) Get(ctx context.Context, input *pb.GetInput) (*pb.Response, error) {
	log.Println("Received request: Get")

	reqUrl, ok := s.storeDependency(input.Url, input.InitialHeader)
	if ok && s.isCircuitOpen(http.MethodGet, reqUrl.Host, reqUrl.Path) {
		return new(pb.Response), status.Error(codes.Aborted, "Request cancelled due to opened circuit")
	}

	res, err := s.doRequest(http.MethodGet, input.Url, nil, input.Header, &s.app.ServiceName, reqUrl)
	if err != nil {
		if errors.Is(err, circuitbreaker.ErrOpenState) {
			return new(pb.Response), status.Error(codes.Aborted, "Request cancelled due to opened circuit")
		}
		return new(pb.Response), status.Error(codes.Internal, "Failed to execute the request")
	}
	defer res.Body.Close()

	return convertResponse(res)
}

func (s HttpServer) Post(ctx context.Context, input *pb.PostInput) (*pb.Response, error) {
	log.Println("Received request: Post")

	reqUrl, ok := s.storeDependency(input.Url, input.InitialHeader)
	if ok && s.isCircuitOpen(http.MethodPost, reqUrl.Host, reqUrl.Path) {
		return new(pb.Response), status.Error(codes.Aborted, "Request cancelled due to opened circuit")
	}

	res, err := s.doRequest(http.MethodPost, input.Url, input.Body, input.Header, &s.app.ServiceName, reqUrl)
	if err != nil {
		if errors.Is(err, circuitbreaker.ErrOpenState) {
			return new(pb.Response), status.Error(codes.Aborted, "Request cancelled due to opened circuit")
		}
		return new(pb.Response), status.Error(codes.Internal, "Failed to execute the request")
	}
	defer res.Body.Close()

	return convertResponse(res)
}

func (s HttpServer) Put(ctx context.Context, input *pb.PutInput) (*pb.Response, error) {
	log.Println("Received request: Put")

	reqUrl, ok := s.storeDependency(input.Url, input.InitialHeader)
	if ok && s.isCircuitOpen(http.MethodPut, reqUrl.Host, reqUrl.Path) {
		return new(pb.Response), status.Error(codes.Aborted, "Request cancelled due to opened circuit")
	}

	res, err := s.doRequest(http.MethodPut, input.Url, input.Body, input.Header, &s.app.ServiceName, reqUrl)
	if err != nil {
		if errors.Is(err, circuitbreaker.ErrOpenState) {
			return new(pb.Response), status.Error(codes.Aborted, "Request cancelled due to opened circuit")
		}
		return new(pb.Response), status.Error(codes.Internal, "Failed to execute the request")
	}
	defer res.Body.Close()

	return convertResponse(res)
}

func (s HttpServer) Delete(ctx context.Context, input *pb.DeleteInput) (*pb.Response, error) {
	log.Println("Received request: Delete")

	reqUrl, ok := s.storeDependency(input.Url, input.InitialHeader)
	if ok && s.isCircuitOpen(http.MethodDelete, reqUrl.Host, reqUrl.Path) {
		return new(pb.Response), status.Error(codes.Aborted, "Request cancelled due to opened circuit")
	}

	res, err := s.doRequest(http.MethodDelete, input.Url, nil, input.Header, &s.app.ServiceName, reqUrl)
	if err != nil {
		if errors.Is(err, circuitbreaker.ErrOpenState) {
			return new(pb.Response), status.Error(codes.Aborted, "Request cancelled due to opened circuit")
		}
		return new(pb.Response), status.Error(codes.Internal, "Failed to execute the request")
	}
	defer res.Body.Close()

	return convertResponse(res)
}
