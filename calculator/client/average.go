package main

import (
	"context"
	"log"
	pb "rpc-demo/calculator/proto"
)

func doAverage(c pb.CalculatorServiceClient) {
	stream, err := c.Average(context.Background())
	if err != nil {
		log.Printf("Average err %v\n", err)
	}

	reqs := []*pb.AverageRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
	}

	for _, req := range reqs {
		log.Printf("Sending req: %d\n", req.Number)
		stream.Send(req)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("doAverage stream.CloseAndRecv error %v\n", err)
	}
	log.Printf("Avg: %f\n", res.Result)
}
