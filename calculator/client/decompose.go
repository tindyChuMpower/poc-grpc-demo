package main

import (
	"context"
	"io"
	"log"
	pb "rpc-demo/calculator/proto"
)

func doDecompose(c pb.CalculatorServiceClient) {
	log.Printf("doDecompose here")
	stream, err := c.Decompose(context.Background(), &pb.DecomposeRequest{
		Number: 66502776,
	})
	if err != nil {
		log.Fatalf("doDecompose error: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("error while reading the stream: %v\n", err)
		}

		log.Printf("%d\n", msg.Result)
	}
}
