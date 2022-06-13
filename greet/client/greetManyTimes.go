package main

import (
	"context"
	"io"
	"log"
	pb "rpc-demo/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	stream, err := c.GreetManyTimes(context.Background(), &pb.GreetRequest{
		FirstName: "First name",
	})
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes")
	}

	for {
		msg, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Printf("GreetManyTimes break ******")
				break
			}
			log.Fatalf("error while reading the stream %v\n", err)
		}
		log.Printf("GreetManyTimes %s\n", msg.Result)
	}

}
