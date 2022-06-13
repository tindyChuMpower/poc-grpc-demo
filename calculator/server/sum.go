package main

import (
	"context"
	"log"
	pb "rpc-demo/calculator/proto"
)

func (*Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)
	return &pb.SumResponse{
		Result: in.A + in.B,
	}, nil
}
