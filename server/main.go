package main

import (
	"fmt"
	"log"
	"net"

	counterproto "github.com/kapibara824/grpc-server_streaming/pb/counter"
	"google.golang.org/grpc"
)

type sever struct{}

func (*sever) Counter(req *counterproto.CounterRequest, stream counterproto.CounterService_CounterServer) error {
	fmt.Printf("Request is %v\n", req.Num)

	number := int(req.GetNum())
	for i := 1; number+1 > i; i++ {
		stream.Send(&counterproto.CounterResponse{
			Result: int64(i),
		})
	}

	return nil

}

func main() {
	fmt.Println("Running gRPC server...")

	lis, err := net.Listen("tcp", "0.0.0.0:8240")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	s := grpc.NewServer()
	counterproto.RegisterCounterServiceServer(s, &sever{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
