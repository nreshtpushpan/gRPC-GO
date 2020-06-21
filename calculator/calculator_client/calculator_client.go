package main

import (
	"context"
	"fmt"
	"github.com/pushpan/gRPC-GO/calculator/calculatorpb"
	"google.golang.org/grpc"
	"io"
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
	doServerStreaming(c)
}

func doUnary(server calculatorpb.CalculatorServiceClient) () {
	fmt.Println("Starting Calculator Unary RPC")
	req := &calculatorpb.SumRequest{
		FirstNumber:  45,
		SecondNumber: 983,
	}
	res, err := server.Sum(context.Background(), req)
	if err != nil {
		log.Fatalln("Error in getting calculator Sum RPC service: %v", err)
	}
	log.Println("Response from calculator service: %v", res)
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a PrimeDecomposition Server Streaming RPC...")
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: 12390392840,
	}
	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling PrimeDecomposition RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}
		fmt.Println(res.GetPrimeFactor())
	}
}
