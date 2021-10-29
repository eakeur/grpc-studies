package main

import (
	"grpc-studies/greet/greetpb"
	"log"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect with server: %v", err)
	}

	defer conn.Close()

	cli := greetpb.NewGreetServiceClient(conn)
	log.Printf("Created client: %v", cli)

}
