package main

import (
	"fmt"
	"github.com/pushpan/gRPC-GO/greet/greetpb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Println("Greet function is invoked %v", req)
	firstName := req.Greeting.GetFirstName()
	result := "Welcome" + firstName
	res := &greetpb.GreetResponse{Result: result}
	return res, nil
}
func main() {
	fmt.Println("This is server")
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalln("Error in listing %v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalln("Error in serving %v", err)
	}
}
