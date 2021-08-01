package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"../hellopb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	fmt.Println("Go client is running")

	certFile := "ssl/ca.crt"

	creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")

	if sslErr != nil {
		log.Fatalf("Failed reading CA certificate %v", sslErr)
	}

	opts := grpc.WithTransportCredentials(creds)

	cc, err := grpc.Dial("localhost:50051", opts) //Con ssl
	// cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) //sin ssl

	if err != nil {
		log.Fatalf("Failed to connect %v", err)
	}

	defer cc.Close()

	c := hellopb.NewHelloServiceClient(cc)

	helloUnary(c)
	// helloServerStreaming(c)
	// goodbyeClientStreaming(c)
	// goodbyeBidiStreaming(c)
}

func helloUnary(c hellopb.HelloServiceClient) {
	fmt.Println("Starting unary RPC Hello")

	req := &hellopb.HelloRequest{
		Hello: &hellopb.Hello{
			FirstName: "Lenin",
			Prefix:    "Joven",
		},
	}

	res, err := c.Hello(context.Background(), req)

	if err != nil {
		log.Fatalf("Error, calling Hello RPC: \n%v", err)
	}

	log.Printf("Response Hello, %v", res.CustomHello)
}

func helloServerStreaming(c hellopb.HelloServiceClient) {
	fmt.Println("Starting server streaming RPC Hello")

	req := &hellopb.HelloManyLanguagesRequest{
		Hello: &hellopb.Hello{
			FirstName: "Lenin",
			Prefix:    "Joven",
		},
	}

	restStream, err := c.HelloManyLanguages(context.Background(), req)
	if err != nil {
		log.Printf("Error calling Hello Many Languages %v", err)
	}

	for {
		msg, err := restStream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading stream %v", err)
		}

		log.Printf("Res from HML: %v\n", msg.GetHelloLanguage())
	}
}

func goodbyeClientStreaming(c hellopb.HelloServiceClient) {
	fmt.Println("Starting goodbye function")

	requests := []*hellopb.HelloGoodbyeRequest{
		&hellopb.HelloGoodbyeRequest{
			Hello: &hellopb.Hello{
				FirstName: "Lenin",
				Prefix:    "Joven",
			},
		},
		&hellopb.HelloGoodbyeRequest{
			Hello: &hellopb.Hello{
				FirstName: "Paul",
				Prefix:    "Menor",
			},
		},
		&hellopb.HelloGoodbyeRequest{
			Hello: &hellopb.Hello{
				FirstName: "Siko",
				Prefix:    "Joven 2",
			},
		},
		&hellopb.HelloGoodbyeRequest{
			Hello: &hellopb.Hello{
				FirstName: "Rodrigo",
				Prefix:    "Joven 3",
			},
		},
		&hellopb.HelloGoodbyeRequest{
			Hello: &hellopb.Hello{
				FirstName: "Ron",
				Prefix:    "Joven 5",
			},
		},
	}

	stream, err := c.HellosGoodbye(context.Background())

	if err != nil {
		log.Printf("Error calling goodbye %v", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)

		stream.Send(req)
		time.Sleep(1000 * time.Microsecond)
	}

	goodbye, err := stream.CloseAndRecv()

	if err != nil {
		log.Printf("Error goodbye receive %v", err)
	}

	fmt.Println(goodbye)
}

func goodbyeBidiStreaming(c hellopb.HelloServiceClient) {
	fmt.Println("Starting goodbye bidi function")

	// Create stream to call server
	stream, err := c.GoodBye(context.Background())
	request := []*hellopb.GoodByeRequest{
		&hellopb.GoodByeRequest{
			Hello: &hellopb.Hello{
				FirstName: "Lenin",
				Prefix:    "Joven",
			},
		},
		&hellopb.GoodByeRequest{
			Hello: &hellopb.Hello{
				FirstName: "Paul",
				Prefix:    "Menor",
			},
		},
		&hellopb.GoodByeRequest{
			Hello: &hellopb.Hello{
				FirstName: "Siko",
				Prefix:    "Joven 2",
			},
		},
		&hellopb.GoodByeRequest{
			Hello: &hellopb.Hello{
				FirstName: "Rodrigo",
				Prefix:    "Joven 3",
			},
		},
		&hellopb.GoodByeRequest{
			Hello: &hellopb.Hello{
				FirstName: "Ron",
				Prefix:    "Joven 5",
			},
		},
	}
	if err != nil {
		log.Printf("Error creating stream %v", err)
	}

	waitc := make(chan struct{})
	// send many messages to the server (go routimes)
	go func() {
		for _, req := range request {
			log.Printf("Sending message %v", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	// Receive messages from the server (go routimes)
	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error receive stream %v", err)
				break
			}
			fmt.Printf("Go it: %v \n", res.GetGoodbye())
		}
		close(waitc)
	}()

	//block when everything is completed or closed
	<-waitc
}
