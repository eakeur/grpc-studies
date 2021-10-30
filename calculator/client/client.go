package main

import (
	"context"
	"grpc-studies/calculator/provider"
	"log"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect with server: %v", err)
	}

	defer conn.Close()

	cli := provider.NewSumServiceClient(conn)

	ret, err := cli.Sum(context.Background(), &provider.SumRequest{X: 1, Y: 1})
	if err != nil {
		log.Fatalf("Failed to greet: %v", err)
	}
	log.Println(ret.Result)

}
