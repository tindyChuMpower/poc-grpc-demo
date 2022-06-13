package main

import (
	"log"
	pb "rpc-demo/calculator/proto"
	"time"
)

func (*Server) Decompose(in *pb.DecomposeRequest, stream pb.CalculatorService_DecomposeServer) error {
	log.Printf("Decompose server here: %v\n", in.Number)
	rest := in.Number
	value := int64(2)
	for {
		if rest == 1 {
			return nil
		}
		if rest%value == 0 {
			time.Sleep(time.Second)
			rest /= value
			stream.Send(&pb.DecomposeResponse{
				Result: value,
			})
		} else {
			value++
		}
	}
}
