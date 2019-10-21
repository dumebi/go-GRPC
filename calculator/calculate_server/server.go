package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/dumebi/grpc-go-course/calculator/calculatepb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Calculate(ctx context.Context, req *calculatepb.CalculateRequest) (*calculatepb.CalculateResponse, error) {
	fmt.Printf("Cal function was invoked with: %v \n", req)
	first := req.GetCal().First
	second := req.GetCal().Second
	result := first + second

	res := &calculatepb.CalculateResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello world!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatepb.RegisterCalculateServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
