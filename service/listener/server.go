package listener

import (
	"log"
	"time"

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
	// Get current time
	now := time.Now().UnixMilli()
	// Validate input
	if input.Expiry <= now {
		return new(empty.Empty), status.Error(codes.InvalidArgument, "Input already expired")
	}
	if len(input.Services) != len(input.Endpoints) {
		return new(empty.Empty), status.Error(codes.InvalidArgument, "Input for services and endpoints are differ in length")
	}
	// Store status
	createdRdServices, createdRdEndpoints, err := s.app.Repositories.Status.CreateFromRdServicesAndRdEndpointsSlice(input.Services, input.Endpoints, "OPEN", input.Expiry)
	if err != nil {
		return new(empty.Empty), status.Error(codes.Internal, "Failed to store open circuits data")
	}
	// Get affected endpoint (only broadcast new created rd endpoints)
	endpoints, err := s.app.Repositories.RequiredService.GetEndpointsByRdServicesAndRdEndpoints(createdRdServices, createdRdEndpoints)
	if err != nil {
		return new(empty.Empty), status.Error(codes.Internal, "Failed to get affected endpoint")
	}
	// Get affected requiring service
	serviceDepMap, endpointDepMap, err := s.app.Repositories.RequiringService.GetDependencyMapByEndpoints(endpoints)
	if err != nil {
		return new(empty.Empty), status.Error(codes.Internal, "Failed to get affected requiring service")
	}
	// Broadcast status
	for serviceAddr := range serviceDepMap {
		err := internal.BroadcastOpenCircuits(serviceAddr, serviceDepMap[serviceAddr], endpointDepMap[serviceAddr], input.Expiry)
		if err != nil {
			log.Printf("OpenCircuits: error on BroadcastOpenCircuits to %s: %v\n", serviceAddr, err)
		}
	}
	return new(empty.Empty), nil
}

func (s ListenerServer) CloseCircuits(ctx context.Context, input *pb.ServiceEndpoints) (*empty.Empty, error) {
	log.Println("Received request: CloseCircuits")
	// Validate input
	if len(input.Services) != len(input.Endpoints) {
		return new(empty.Empty), status.Error(codes.InvalidArgument, "Input for services and endpoints are differ in length")
	}
	// Store status
	deletedRdServices, deletedRdEndpoints, err := s.app.Repositories.Status.DeleteWhereRdServicesAndRdEndpointsInSlice(input.Services, input.Endpoints)
	if err != nil {
		return new(empty.Empty), status.Error(codes.Internal, "Failed to remove open circuits data")
	}
	// Get affected endpoint (only broadcast new deleted rd endpoints)
	endpoints, err := s.app.Repositories.RequiredService.GetEndpointsByRdServicesAndRdEndpoints(deletedRdServices, deletedRdEndpoints)
	if err != nil {
		return new(empty.Empty), status.Error(codes.Internal, "Failed to get affected endpoint")
	}
	// Get affected requiring service
	serviceDepMap, endpointDepMap, err := s.app.Repositories.RequiringService.GetDependencyMapByEndpoints(endpoints)
	if err != nil {
		return new(empty.Empty), status.Error(codes.Internal, "Failed to get affected requiring service")
	}
	// Broadcast status
	for serviceAddr := range serviceDepMap {
		err := internal.BroadcastCloseCircuits(serviceAddr, serviceDepMap[serviceAddr], endpointDepMap[serviceAddr])
		if err != nil {
			log.Printf("CloseCircuits: error on BroadcastCloseCircuits to %s: %v\n", serviceAddr, err)
		}
	}
	return new(empty.Empty), nil
}
