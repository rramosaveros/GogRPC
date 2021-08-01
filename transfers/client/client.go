package main

import (
	"context"
	"fmt"
	"log"

	"../transferspb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("Go client is running")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect %v", err)
	}

	defer cc.Close()

	c := transferspb.NewTransfersServiceClient(cc)

	transfer(c)
}

func transfer(c transferspb.TransfersServiceClient) {
	fmt.Println("starting unary PRC Transfer")

	req := &transferspb.TransfersRequest{
		OriginAccount: "87655678909876",
		// OriginAccount:   "87655678909875", //para probar el error
		ReceiverAccount: "98555682009402",
		// Amount:          100,
		Amount: 0, //Para probar el error
	}

	res, err := c.Transfer(context.Background(), req)

	if err != nil {
		respErr, exception := status.FromError(err)
		if exception {
			fmt.Printf("---Exception message from server: %v\n", respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("Sorry the amount has to be more than 0")
				return
			}
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("Parece que escribiste mal el número de cuenta de origen")
				return
			}
		} else {
			log.Fatalf("Big Error calling Transfer: %v", err)
			return
		}
	}

	log.Printf("Response Transfer: %v", res.GetOperationDate())
}
