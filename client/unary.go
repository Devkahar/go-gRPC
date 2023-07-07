package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Devkahar/go-gRPC/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancle := context.WithTimeout(context.Background(), time.Second)

	defer cancle()
	res, err := client.SayHello(ctx, &pb.NoParams{})

	if err != nil {
		log.Fatalf("could not greet %v", err)
	}

	log.Printf("%s", res.Message)

}
