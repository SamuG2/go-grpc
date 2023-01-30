package main

import (
	"log"

	"github.com/SamuG2/go-grpc/client/greet"

	pb "github.com/SamuG2/go-grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("couldn't connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Jesús", "María", "José"},
	}

	greet.CallSayHello(client)
	greet.CallSayHelloServerStream(client, names)
	greet.CallSayHelloClientStream(client, names)
	greet.CallSayHelloBidirectionalStream(client, names)

}
