package main

import (
	"context"
	"log"
	pb "rpc-demo/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doCalculator was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		A: 1,
		B: 2,
	})
	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}

	log.Printf("Calculator: %s\n", res)
}
