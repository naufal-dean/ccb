package listener

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/naufal-dean/ccb/app"
	pb "github.com/naufal-dean/ccb/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
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
	s.app.Repositories.Status.CreateFromOneServiceAndManyEndpoints(input.Service, input.Endpoints, "OPEN")
	return new(empty.Empty), nil
}

func (s ListenerServer) CloseCircuits(ctx context.Context, input *pb.ServiceEndpoints) (*empty.Empty, error) {
	log.Println("Received request: CloseCircuits")
	s.app.Repositories.Status.DeleteWhereOneServiceAndManyEndpointsEqual(input.Service, input.Endpoints)
	return new(empty.Empty), nil
}
