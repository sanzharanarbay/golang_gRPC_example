package main

import (
	"fmt"
	"golang-gRPC-example/api"
	"log"
	"net"

	"google.golang.org/grpc"
	"golang-gRPC-example/server"
)

func main() {
	fmt.Println("Go gRPC starts!")
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := server.Server{}

	grpcServer := grpc.NewServer()

	api.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
