package main

import (
	"log"

	pb "github.com/Devkahar/go-gRPC/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":5000"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect %v", err)
	}

	defer conn.Close()
	client := pb.NewGreetServiceClient(conn)

	names := &pb.NameList{
		Names: []string{"Dev", "Alic", "Bob"},
	}
	// callSayHello(client)
	// callSayHelloServerStreaming(client, names)
	// callSayHelloClientStreamServer(client, names)
	callSayHelloBidirectionalStream(client, names)

}
