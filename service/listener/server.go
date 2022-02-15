package listener

import (
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/naufal-dean/ccb/app"
	pb "github.com/naufal-dean/ccb/protobuf"
	"github.com/naufal-dean/ccb/service/internal"
)

type ListenerServer struct {
	app app.App
	pb.UnimplementedListenerServer
}

func New(app app.App) *ListenerServer {
	return &ListenerServer{app: app}
}

func (s *ListenerServer) Register(sr grpc.ServiceRegistrar) {
	pb.RegisterListenerServer(sr, s)
}

func (s ListenerServer) OpenCircuits(ctx context.Context, input *pb.ServiceEndpoints) (*empty.Empty, error) {
	log.Println("Received request: OpenCircuits")
	// Store status
	// TODO: only broadcast new updated status, to prevent circular broadcast
	err := s.app.Repositories.Status.CreateFromOneRdServiceAndManyRdEndpoints(input.Service, input.Endpoints, "OPEN")
	if err != nil {
		return new(empty.Empty), status.Error(codes.Internal, "Failed to store open circuits data")
	}
	// Get affected endpoint
	endpoints, err := s.app.Repositories.RequiredService.GetEndpointsByRdServiceAndRdEndpoints(input.Service, input.Endpoints)
	if err != nil {
		return new(empty.Empty), status.Error(codes.Internal, "Failed to get affected endpoint")
	}
	log.Printf("OpenCircuits: Endpoints: %v\n", endpoints)
	// Get affected requiring service
	dependencyMap, err := s.app.Repositories.RequiringService.GetDependencyMapByEndpoints(endpoints)
	if err != nil {
		return new(empty.Empty), status.Error(codes.Internal, "Failed to get affected requiring service")
	}
	log.Printf("OpenCircuits: Dep Map: %v\n", dependencyMap)
	// Broadcast status
	for serviceAddr, endpoints := range dependencyMap {
		log.Printf("OpenCircuits: Broadcast: %v, %v, %v\n", s.app.ServiceName, serviceAddr, endpoints)
		internal.BroadcastOpenCircuits(s.app.ServiceName, serviceAddr, endpoints)
	}
	return new(empty.Empty), nil
}

func (s ListenerServer) CloseCircuits(ctx context.Context, input *pb.ServiceEndpoints) (*empty.Empty, error) {
	log.Println("Received request: CloseCircuits")
	// Store status
	// TODO: only broadcast new updated status, to prevent circular broadcast
	err := s.app.Repositories.Status.DeleteWhereOneRdServiceAndManyRdEndpointsEqual(input.Service, input.Endpoints)
	if err != nil {
		return new(empty.Empty), status.Error(codes.Internal, "Failed to remove open circuits data")
	}
	// Get affected endpoint
	endpoints, err := s.app.Repositories.RequiredService.GetEndpointsByRdServiceAndRdEndpoints(input.Service, input.Endpoints)
	if err != nil {
		return new(empty.Empty), status.Error(codes.Internal, "Failed to get affected endpoint")
	}
	// Get affected requiring service
	dependencyMap, err := s.app.Repositories.RequiringService.GetDependencyMapByEndpoints(endpoints)
	if err != nil {
		return new(empty.Empty), status.Error(codes.Internal, "Failed to get affected requiring service")
	}
	// Broadcast status
	for serviceAddr, endpoints := range dependencyMap {
		internal.BroadcastCloseCircuits(s.app.ServiceName, serviceAddr, endpoints)
	}
	return new(empty.Empty), nil
}
