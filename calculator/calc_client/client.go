package main

import (
	"context"
	"fmt"
	"grpc-go-course/calculator/calculatorpb"
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
	doUnary(c)

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
