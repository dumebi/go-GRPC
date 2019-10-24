package main

import (
	"context"
	"fmt"
	"io"
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

func (*server) ComputeAverage(stream calculatepb.CalculateService_ComputeAverageServer) error {
	fmt.Printf("Compute Average function was invoked \n")
	numbers := 0
	count := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// we have finished the stream
			average := numbers / count
			return stream.SendAndClose(&calculatepb.ComputeAverageResponse{
				Result: int32(average),
			})
		}
		if err != nil {
			log.Fatalf("Failed to read client stream: %v", err)
		}

		number := req.GetCa().Number
		numbers += int(number)
		count++
	}
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
