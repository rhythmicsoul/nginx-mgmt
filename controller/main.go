package main

import (
	"context"
	"log"
	"net"

	"github.com/rhythmicsoul/nginx-mgmt/proto/controller"
	"google.golang.org/grpc"
)

type server struct {
	controller.UnimplementedAddServiceServer
}

func (s *server) NewAgentToken(ctx context.Context, req *controller.Empty) (*controller.AgentToken, error) {
	return &controller.AgentToken{Token: "an agent token"}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:9091")
	if err != nil {
		log.Fatalf("Starting of server failed: %v", err)
	}

	s := grpc.NewServer()
	controller.RegisterAddServiceServer(s, &server{})
	log.Printf("server listening at: %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
