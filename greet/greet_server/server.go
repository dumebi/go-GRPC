package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/dumebi/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with: %v \n", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "hello " + firstName

	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("Greet many times function was invoked with: %v \n", req)
	firstName := req.GetGreeting().GetFirstName()

	for i := 0; i < 10; i++ {
		result := "hello " + firstName + " number " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesresponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Printf("Long greet function was invoked \n")
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// we have finished the stream
			return stream.SendAndClose(&greetpb.LongGreetresponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Failed to read client stream: %v", err)
		}

		firstName := req.GetGreeting().GetFirstName()
		result += "Hello "+ firstName + "! "
	}
}

func main() {
	fmt.Println("Hello world!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
