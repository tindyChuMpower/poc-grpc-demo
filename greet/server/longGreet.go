package main

import (
	"fmt"
	"io"
	"log"
	pb "rpc-demo/greet/proto"
)

func (*Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Printf("LongGreet here")

	res := ""

	for {
		result, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return stream.SendAndClose(&pb.GreetResponse{
					Result: res,
				})
			}

			log.Fatalf("LongGreet err %v\n", err)
		}

		res += fmt.Sprintf("Hello %s\n", result.FirstName)
	}
}
