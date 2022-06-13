package main

import (
	"io"
	"log"
	pb "rpc-demo/calculator/proto"
	"time"
)

func (*Server) Average(stream pb.CalculatorService_AverageServer) error {
	log.Printf("Average here")
	var res int32 = 0
	count := 0

	for {
		log.Printf("Calculating")
		result, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Printf("Finish")
				return stream.SendAndClose(&pb.AverageResponse{
					Result: float64(res) / float64(count),
				})
			}

			log.Fatalf("LongGreet err %v\n", err)
		}
		res += result.Number
		count++

		time.Sleep(time.Second)
	}
}
