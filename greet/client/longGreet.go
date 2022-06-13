package main

import (
	"context"
	"log"
	pb "rpc-demo/greet/proto"
	"time"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Printf("doLongGreet here")

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Printf("doLongGreet err %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Name 1"},
		{FirstName: "Name 2"},
		{FirstName: "Name 3"},
	}

	for _, req := range reqs {
		log.Printf("Sending req: %s\n", req.FirstName)
		stream.Send(req)
		time.Sleep(time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("doLongGreet stream.CloseAndRecv error %v\n", err)
	}
	log.Printf(res.Result)
}
