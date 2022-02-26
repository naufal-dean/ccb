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
	createdRdEndpoints, err := s.app.Repositories.Status.CreateFromOneRdServiceAndManyRdEndpoints(input.Service, input.Endpoints, "OPEN")
	if err != nil {
		return new(empty.Empty), status.Error(codes.Internal, "Failed to store open circuits data")
	}
	// Get affected endpoint (only broadcast new created rd endpoints)
	endpoints, err := s.app.Repositories.RequiredService.GetEndpointsByRdServiceAndRdEndpoints(input.Service, createdRdEndpoints)
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
		err := internal.BroadcastOpenCircuits(s.app.ServiceName, serviceAddr, endpoints)
		if err != nil {
			log.Printf("OpenCircuits: error on BroadcastOpenCircuits to %s: %v\n", serviceAddr, err)
		}
	}
	return new(empty.Empty), nil
}

func (s ListenerServer) CloseCircuits(ctx context.Context, input *pb.ServiceEndpoints) (*empty.Empty, error) {
	log.Println("Received request: CloseCircuits")
	// Store status
	deletedRdEndpoints, err := s.app.Repositories.Status.DeleteWhereOneRdServiceAndManyRdEndpointsEqual(input.Service, input.Endpoints)
	if err != nil {
		return new(empty.Empty), status.Error(codes.Internal, "Failed to remove open circuits data")
	}
	// Get affected endpoint (only broadcast new deleted rd endpoints)
	endpoints, err := s.app.Repositories.RequiredService.GetEndpointsByRdServiceAndRdEndpoints(input.Service, deletedRdEndpoints)
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
		err := internal.BroadcastCloseCircuits(s.app.ServiceName, serviceAddr, endpoints)
		if err != nil {
			log.Printf("CloseCircuits: error on BroadcastCloseCircuits to %s: %v\n", serviceAddr, err)
		}
	}
	return new(empty.Empty), nil
}
