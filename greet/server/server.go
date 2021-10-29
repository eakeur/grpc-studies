package main

import (
	"context"
	"fmt"
	"grpc-studies/greet/greetpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
}

func (s *server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was called with name %v %v", req.Greeting.FirstName, req.Greeting.LastName)
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()
	return &greetpb.GreetResponse{
		Result: "Hello, " + firstName + " " + lastName,
	}, nil
}

func main() {

	// 50051 is the gRPC default port
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Could not listen to port 50051: %v", err)
	}

	srv := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(srv, &server{})

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}
}
