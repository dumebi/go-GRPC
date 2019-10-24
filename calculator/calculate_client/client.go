package main

import (
	"context"
	"fmt"
	"log"
	"time"

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
	// doUnary(c)
	doClientStreaming(c)
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

func doClientStreaming(c calculatepb.CalculateServiceClient) {
	fmt.Println("Starting to do a client stream RPC...")
	requests := []*calculatepb.ComputeAverageRequest{
		&calculatepb.ComputeAverageRequest{
			Ca: &calculatepb.ComputeAverage{
				Number: 5,
			},
		},
		&calculatepb.ComputeAverageRequest{
			Ca: &calculatepb.ComputeAverage{
				Number: 5,
			},
		},
		&calculatepb.ComputeAverageRequest{
			Ca: &calculatepb.ComputeAverage{
				Number: 5,
			},
		},
		&calculatepb.ComputeAverageRequest{
			Ca: &calculatepb.ComputeAverage{
				Number: 5,
			},
		},
		&calculatepb.ComputeAverageRequest{
			Ca: &calculatepb.ComputeAverage{
				Number: 5,
			},
		},
		&calculatepb.ComputeAverageRequest{
			Ca: &calculatepb.ComputeAverage{
				Number: 5,
			},
		},
		&calculatepb.ComputeAverageRequest{
			Ca: &calculatepb.ComputeAverage{
				Number: 5,
			},
		},
		&calculatepb.ComputeAverageRequest{
			Ca: &calculatepb.ComputeAverage{
				Number: 5,
			},
		},
	}
	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Compute Average RPC: %v", err)
	}

	// iterate over our slice and send
	for _, req := range requests {
		log.Printf("sending request: %v\n", req)
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from Compute Average RPC: %v", err)
	}
	fmt.Printf("Compute Average Response: %v\n", response)
}
