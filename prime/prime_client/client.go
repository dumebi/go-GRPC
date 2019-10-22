package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/dumebi/grpc-go-course/prime/primepb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, I'm a client!")
	conn, err := grpc.Dial("localhost: 50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer conn.Close()
	c := primepb.NewPrimeNumberServiceClient(conn)
	// fmt.Printf("created client: %f", c)
	// doUnary(c)
	doServerStreaming(c)
}

func doServerStreaming(c primepb.PrimeNumberServiceClient) {
	fmt.Println("Starting to do a streaming RPC...")
	req := &primepb.PrimeNumberManyTimesRequest{
		Number: &primepb.PrimeNumber{
			Number: "960",
		},
	}
	responseStream, err := c.PrimeNumberManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling server stream RPC: %v", err)
	}

	for {
		msg, err := responseStream.Recv()
		if err == io.EOF {
			// we reached end of stream
			break
		}
		if err != nil {
			log.Fatalf("Error while reading server stream RPC: %v", err)
		}
		log.Printf("Response from prime number many times : %v", msg.GetResult())
	}
}
