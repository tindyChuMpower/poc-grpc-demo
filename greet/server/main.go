package main

import (
	"fmt"
	pb "rpc-demo/greet/proto"
	"time"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

// func main() {
// 	lis, err := net.Listen("tcp", addr)
// 	if err != nil {
// 		log.Fatalf("Failed to listen on: %v\n", err)
// 	}

// 	log.Printf("listening on %s\n", addr)

// 	s := grpc.NewServer()
// 	pb.RegisterGreetServiceServer(s, &Server{})

// 	if err = s.Serve(lis); err != nil {
// 		log.Fatalf("Failed tosServe: %v\n", err)
// 	}

// }

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done

	fmt.Println("done")
}
