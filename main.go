package main

import (
	"context"
	"google.golang.org/grpc"
	"net"
	"time"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()

	sleepService := SleepService{}

	RegisterSleepServiceServer(server, &sleepService)

	panic(server.Serve(l))
}

type SleepService struct{}

func (s *SleepService) SleepFor(ctx context.Context, req *SleepRequest) (*SleepResponse, error) {
	time.Sleep(time.Duration(req.Milliseconds) * time.Millisecond)
	return &SleepResponse{}, nil
}
