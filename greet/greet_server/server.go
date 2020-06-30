package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc-go-course-master/greet/greetpb"

	"log"
	"net"
)

type server struct {

}
func main() {
	fmt.Println(2)

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Server(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
