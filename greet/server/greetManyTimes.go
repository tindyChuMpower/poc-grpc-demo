package main

import (
	"fmt"
	"log"
	pb "rpc-demo/greet/proto"
	"time"
)

func (*Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked with: %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", in.FirstName, i)
		stream.Send(&pb.GreetResponse{
			Result: res,
		})
		time.Sleep(time.Second)
	}
	return nil
}
