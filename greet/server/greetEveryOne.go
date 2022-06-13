package main

import (
	"io"
	"log"
	pb "rpc-demo/greet/proto"
)

func (*Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Printf("GreetEveryone here")

	for {
		result, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}

			log.Fatalf("LongGreet err %v\n", err)
		}

		res := "Hello " + result.FirstName
		if err := stream.Send(&pb.GreetResponse{Result: res}); err != nil {
			log.Fatalf("LongGreet err %v\n", err)
		}
	}
}
