package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/dumebi/grpc-go-course/prime/primepb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) PrimeNumberManyTimes(req *primepb.PrimeNumberManyTimesRequest, stream primepb.PrimeNumberService_PrimeNumberManyTimesServer) error {
	fmt.Printf("Greet many times function was invoked with: %v \n", req)
	number, err := strconv.Atoi(req.GetNumber().Number)
	if err != nil {
		log.Fatalf("Failed to convert string to number: %v", err)
	}

	k := 2
	N := number
	for N > 1 {
		if N%k == 0 {
			result := strconv.Itoa(k) + ","
			res := &primepb.PrimeNumberManyTimesresponse{
				Result: result,
			}
			N = N / k
			stream.Send(res)
			time.Sleep(1000 * time.Millisecond)
		} else {
			k = k + 1
		}
	}
	return nil
}

func main() {
	fmt.Println("Hello world!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	primepb.RegisterPrimeNumberServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
