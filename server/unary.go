package main

import (
	"context"

	pb "github.com/Devkahar/go-gRPC/proto"
)

func (c *helloServer) SayHello(ctx context.Context, req *pb.NoParams) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
