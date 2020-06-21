package main

import (
	"context"
	"fmt"
	"github.com/pushpan/gRPC-GO/greet/greetpb"
	"google.golang.org/grpc"
	"log"
)

func main() {

	fmt.Println("Greet Client Side")
	cc, err := grpc.Dial("localhost:50051")
	if err != nil {
		log.Fatalln("Error in Dailing %v", err)
	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	doUnary(c)
}

func doUnary(server greetpb.GreetServiceClient) () {
	fmt.Println("Starting Unary RPC")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{FirstName: "Pushpan", LastName: "Nresht"},
	}
	res, err := server.Greet(context.Background(), req)
	if err != nil {
		log.Fatalln("Error in getting greet RPC service: %v", err)
	}
	log.Println("Response from greet service: %v", res)
}
