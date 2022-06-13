package main

import (
	"context"
	"io"
	"log"
	pb "rpc-demo/greet/proto"
	"time"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone here")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("GreetEveryone err: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Name a"},
		{FirstName: "Name b"},
		{FirstName: "Name c"},
	}

	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v\n", req)
			stream.Send(req)
			time.Sleep(time.Second)
		}
		stream.CloseSend()
	}()

	waitc := make(chan struct{})
	go func() {
		for {
			res, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					log.Printf("doGreetEveryone break ******")
					break
				}
				log.Fatalf("error while reading the stream %v\n", err)
			}

			log.Printf("Received: %v\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
