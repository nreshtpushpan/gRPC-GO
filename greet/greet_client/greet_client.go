package main

import (
	"fmt"
	"github.com/pushpan/gRPC-GO/greet/greetpb"
	"google.golang.org/grpc"
	"log"
)

func main() {

	fmt.Println("Client Side")
	cc,err := grpc.Dial("localhost:50051")
	if err != nil {
		log.Fatalln("Error in Dailing %v", err)
	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	fmt.Println(c)

}
