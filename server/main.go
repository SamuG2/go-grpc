package main

import (
	"log"
	"net"

	pb "github.com/SamuG2/go-grpc/proto"
	"github.com/SamuG2/go-grpc/server/greet"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

// type helloServer struct {
// 	pb.GreetServiceServer
// }

// func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
// 	fmt.Println("Say Hello")
// 	return &pb.HelloResponse{
// 		Message: "Hello",
// 	}, nil
// }

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
	//create new gRPC server
	grpcServer := grpc.NewServer()
	// register greet service
	pb.RegisterGreetServiceServer(grpcServer, &greet.HelloServer{})
	log.Printf("server started at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}

}
