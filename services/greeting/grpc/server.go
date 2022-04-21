package grpc

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "github.com/knwoop/microsercices-example/services/greeting/proto"
)

var _ pb.GreeterServer = (*server)(nil)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func New() (*grpc.Server, error) {
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	return s, nil
}