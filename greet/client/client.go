package main

import (
	"context"
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

	req := &greetpb.Greeting{FirstName: "Igor", LastName: "Reis"}

	for i := 0; i < 100000; i++ {
		ret, err := cli.Greet(context.Background(), &greetpb.GreetRequest{Greeting: req})
		if err != nil {
			log.Fatalf("Failed to greet: %v", err)
		}
		log.Println(ret.Result)
	}

}
