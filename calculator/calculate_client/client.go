package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dumebi/grpc-go-course/calculator/calculatepb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, I'm a client!")
	conn, err := grpc.Dial("localhost: 50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer conn.Close()
	c := calculatepb.NewCalculateServiceClient(conn)
	// fmt.Printf("created client: %f", c)
	doUnary(c)
}

func doUnary(c calculatepb.CalculateServiceClient) {
	fmt.Println("Starting to do a unary RPC...")
	req := &calculatepb.CalculateRequest{
		Cal: &calculatepb.Calculate{
			First:  10,
			Second: 3,
		},
	}
	response, err := c.Calculate(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from greet: %v", response.Result)
}
