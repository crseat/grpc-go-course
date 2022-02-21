package main

import (
	"context"
	"fmt"
	"grpc-go-course/calculator/calculatorpb"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("Hello I'm a client")
	creds := insecure.NewCredentials()
	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalcServiceClient(cc)
	// fmt.Printf("Created client: %f", c)
	//doUnary(c)
	doServerStreaming(c)

}

func doUnary(c calculatorpb.CalcServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &calculatorpb.CalcRequest{
		Nums: &calculatorpb.Nums{
			FirstNum:  3,
			SecondNum: 5,
		},
	}

	res, err := c.Calculate(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Calculate RPC: %v", err)
	}
	log.Printf("Response from Calculate: %v", res.Result)

}

func doServerStreaming(c calculatorpb.CalcServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC...")
	req := &calculatorpb.PrimeNumDecompRequest{
		Num: 120,
	}
	resStream, err := c.PrimeNumDecomp(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling PrimeNumDecomp RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// stream has been closed
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response from doServerStreaming: %v", msg.GetResult())
	}
}
