package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/dumebi/grpc-go-course/greet/greetpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, I'm a client!")
	conn, err := grpc.Dial("localhost: 50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer conn.Close()
	c := greetpb.NewGreetServiceClient(conn)
	// fmt.Printf("created client: %f", c)
	// doUnary(c)
	doServerStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Jude",
			LastName:  "Dike",
		},
	}
	response, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from greet: %v", response.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a streaming RPC...")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Jude",
			LastName:  "Dike",
		},
	}
	responseStream, err := c.GreetManyTimes(context.Background(), req)
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
		log.Printf("Response from greet many times : %v", msg.GetResult())
	}
}
