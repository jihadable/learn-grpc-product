package main

import (
	"log"
	"net"

	"github.com/jihadable/learn-grpc-product/service"
	productPb "github.com/jihadable/learn-grpc-proto/product"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	productService := service.NewProductService()
	productPb.RegisterProductServiceServer(grpcServer, productService)

	log.Println("Category gRPC server running on :50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
