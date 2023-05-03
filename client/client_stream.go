package main

import (
	"context"
	pb "grpc-demo/proto"
	"log"
	"time"
)

func callSayHelloClientStream(cli pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Client streaming started")

	stream, err := cli.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send name: %e", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending: %e", err)
		}
		log.Printf("Sent the request with name: %e", err)

		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("Client streaming finished")
	if err != nil {
		log.Fatalf("Error while receiving")
	}

	log.Printf("%v", res)
}
