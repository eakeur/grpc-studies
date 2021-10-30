package main

import (
	"context"
	"fmt"
	"grpc-studies/calculator/provider"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
}

func (s *server) Sum(ctx context.Context, req *provider.SumRequest) (*provider.SumResponse, error) {
	fmt.Printf("Greet function was called with values %v %v", req.X, req.Y)
	return &provider.SumResponse{
		Result: req.X + req.Y,
	}, nil
}

func main() {

	// 50051 is the gRPC default port
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Could not listen to port 50051: %v", err)
	}

	srv := grpc.NewServer()

	provider.RegisterSumServiceServer(srv, &server{})

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}
}
