package main

import (
	"fmt"
	"github.com/pushpan/gRPC-GO/greet/greetpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {}

func main() {

	fmt.Println("This is server")
	lis,err := net.Listen("tcp","localhost:50051")
	if err != nil {
		log.Fatalln("Error in listing %v",err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s,server{})
	if err := s.Serve(lis) ; err != nil {
		log.Fatalln("Error in serving %v",err)
	}

}
