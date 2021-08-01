package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"../hellopb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct{}

func (*server) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	fmt.Println("hello function was called with %v \n", req)

	firstName := req.GetHello().GetFirstName()
	prefix := req.GetHello().GetPrefix()

	customHello := "Welcome: " + prefix + " " + firstName

	res := &hellopb.HelloResponse{
		CustomHello: customHello,
	}
	return res, nil
}

func (*server) HelloManyLanguages(req *hellopb.HelloManyLanguagesRequest, stream hellopb.HelloService_HelloManyLanguagesServer) error {
	fmt.Println("Hello Many times function was invoked with %v\n", req)

	langs := [10]string{"Salut: ", "Hello! ", "Ni hao: ", "Aló! ", "привет! ", "Schalom! ", "Hola! ", "Yassou! ", "Nej! ", "Konnichiwa! "}

	firstName := req.GetHello().GetFirstName()
	prefix := req.GetHello().GetPrefix()

	for _, helloLang := range langs {
		helloLanguage := helloLang + prefix + " " + firstName

		res := &hellopb.HelloManyLanguagesResponse{
			HelloLanguage: helloLanguage,
		}

		stream.Send(res)

		time.Sleep(1000 * time.Millisecond) //one second
	}
	return nil
}

func (*server) HellosGoodbye(stream hellopb.HelloService_HellosGoodbyeServer) error {
	fmt.Println("Goodbye function was invoked")

	goodbye := "Goodbye guys: "

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			// Once is finished the stream we gonna send the response
			return stream.SendAndClose(&hellopb.HelloGoodbyeResponse{
				Goodbye: goodbye,
			})
		}

		if err != nil {
			log.Fatalf("Error reading the client stream %v", err)
		}

		firstName := req.GetHello().GetFirstName()
		prefix := req.GetHello().GetPrefix()

		goodbye += prefix + " " + firstName + " , "
	}
}

func (*server) GoodBye(stream hellopb.HelloService_GoodByeServer) error {
	fmt.Println("Goodbye bidirectional function was invoked")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error reading the client stream %v", err)
			return nil
		}

		firstName := req.GetHello().GetFirstName()
		prefix := req.GetHello().GetPrefix()

		goodBye := "Goodbye " + prefix + " " + firstName + " :( "

		sendErr := stream.Send(&hellopb.GoodByeResponse{
			Goodbye: goodBye,
		})

		if sendErr != nil {
			log.Fatalf("Error sending to the client %v", sendErr)
		}
	}
}

func main() {
	fmt.Println("Hello, Go Server is running")

	certFile := "ssl/server.crt"
	keyFile := "ssl/server.pem"

	//linux
	// certFile := "../../ssl/server.crt"
	// keyFile := "../../ssl/server.pem"

	creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)

	if sslErr != nil {
		log.Fatalf("Failed loading certificates %v", sslErr)
	}

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to liste %v", err)
	}

	opts := grpc.Creds(creds)
	s := grpc.NewServer(opts)

	hellopb.RegisterHelloServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
