package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/graphql-services/oauth/grpc"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	log.Printf("request for user: %s, scopes: %s\n", req.UserID, req.Scopes)
	return &pb.ValidateResponse{Valid: true, Scopes: req.Scopes}, nil
}

const defaultPort = "80"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterScopeValidatorServer(s, &server{})

	fmt.Println("Listening on", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
