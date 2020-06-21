package main

import (
	"context"
	"fmt"
	"github.com/pushpan/gRPC-GO/greet/greetpb"
	"google.golang.org/grpc"
	"io"
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
	doServerStreaming(c)
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

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Stephane",
			LastName:  "Maarek",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTimes RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}

}
