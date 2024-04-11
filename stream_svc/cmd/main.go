package main

import (
	"fmt"
	"log"
	"net"
	"stream_svc/config"
	"stream_svc/internal/pb"
	"stream_svc/internal/services"

	"google.golang.org/grpc"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	lis, err := net.Listen("tcp", config.Port)
	if err != nil {
		fmt.Println("-------", err)
	}
	fmt.Println("Stream Service on:", config.Port)

	s := services.StreamServer{}

	grpcServer := grpc.NewServer()
	pb.RegisterStreamServiceServer(grpcServer, &s)

	fmt.Println("Starting gRPC server...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
