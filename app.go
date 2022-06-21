package main

import (
	"context"
	"fmt"
	pb "gitHub.com/apigee/dummy-grpc/greeting"
	"log"
	"net"

	"google.golang.org/grpc"
)

type greetingServiceServer struct {
	pb.UnimplementedGreetingServiceServer
}

func (s *greetingServiceServer) SayHello(ctx context.Context, message *pb.Message) (*pb.Message, error) {
	return &pb.Message{Body: fmt.Sprintf("Hello from the %+v", message.Body)}, nil
}

func newServer() *greetingServiceServer {
	s := &greetingServiceServer{}
	return s
}

func main() {
	listner, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterGreetingServiceServer(grpcServer, newServer())
	if err := grpcServer.Serve(listner); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
