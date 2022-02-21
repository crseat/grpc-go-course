package main

import (
	"context"
	"fmt"
	"grpc-go-course/calculator/calculatorpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	calculatorpb.UnimplementedCalcServiceServer
}

func (*server) Calculate(ctx context.Context, req *calculatorpb.CalcRequest) (*calculatorpb.CalcResponse, error) {
	fmt.Printf("Calculate function was invoked with %v\n", req)
	firstNum := req.GetNums().GetFirstNum()
	secondNum := req.GetNums().GetSecondNum()
	result := firstNum + secondNum
	res := calculatorpb.CalcResponse{
		Result: result,
	}
	return &res, nil
}

func (*server) PrimeNumDecomp(req *calculatorpb.PrimeNumDecompRequest, stream calculatorpb.CalcService_PrimeNumDecompServer) error {
	fmt.Printf("PrimeNumDecomp function was invoked with %v\n", req)
	num := req.GetNum()
	var k int32 = 2
	for num > 1 {
		if num%k == 0 {
			res := &calculatorpb.PrimeNumDecomResponse{
				Result: k,
			}
			stream.Send(res)
			num = num / k
		} else {
			k = k + 1
		}
		//time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalcServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to sere: %v", err)
	}
}
