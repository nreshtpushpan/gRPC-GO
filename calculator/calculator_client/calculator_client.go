package main

import (
	"context"
	"fmt"
	"github.com/pushpan/gRPC-GO/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
)

func main() {

	fmt.Println("Calculator Client Side")
	cc, err := grpc.Dial("localhost:50051")
	if err != nil {
		log.Fatalln("Error in Dailing %v", err)
	}
	defer cc.Close()
	c := calculatorpb.NewCalculatorServiceClient(cc)
	doUnary(c)
}

func doUnary(server calculatorpb.CalculatorServiceClient) () {
	fmt.Println("Starting Calculator Unary RPC")
	req := &calculatorpb.SumRequest{
		FirstNumber:45,
		SecondNumber:983,
	}
	res, err := server.Sum(context.Background(), req)
	if err != nil {
		log.Fatalln("Error in getting calculator Sum RPC service: %v", err)
	}
	log.Println("Response from calculator service: %v", res)
}
