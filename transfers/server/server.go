package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"../transferspb"

	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

// The server choose when how many responses send for the client

func (*server) Transfer(ctx context.Context, req *transferspb.TransfersRequest) (*transferspb.TransfersResponse, error) {
	fmt.Println("Stating transfer", req)

	originAccount := req.GetOriginAccount()
	//receiverAccount := req.GetReceiverAccount()

	amount := req.GetAmount()

	if amount <= 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Error amount less or equal to 0 with amount $ %v ", amount),
		)
	}

	if originAccount != "87655678909876" {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Esta cuenta no pertenece al banco: %v", originAccount),
		)
	}

	// logic to validate accounts number
	// transfer logic code

	year, month, day := time.Now().Date()

	res := &transferspb.TransfersResponse{
		OperationDate: &date.Date{
			Year:  int32(year),
			Month: int32(month),
			Day:   int32(day),
		},
	}

	return res, nil
}

func main() {
	fmt.Println("Transfers, Go Server is running")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to liste %v", err)
	}

	s := grpc.NewServer()

	transferspb.RegisterTransfersServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
