package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Devkahar/go-gRPC/proto"
)

func callSayHelloClientStreamServer(client pb.GreetServiceClient, names *pb.NameList) error {
	log.Println("Client Streaming Stated")
	stream, err := client.SayHelloClientStreaming(context.Background())

	if err != nil {
		log.Fatalf("Could not send names %v", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			log.Printf("Error while sending %v", err)
		}
		log.Printf("Send request with name %v", name)
		time.Sleep(2 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	log.Println("Client Streaming end")

	if err != nil {
		log.Fatalf("Error while reciving %v", err)
	}
	log.Printf("Recive from server %v", res.Messages)
	return nil
}
