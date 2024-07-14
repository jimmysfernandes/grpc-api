package main

import (
	"fmt"
	"log"
	"net"

	"github.com/jimmysfernandes/grpc-api/customer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/jimmysfernandes/grpc-api/proto"
)

func main() {
	creds, err := credentials.NewServerTLSFromFile("server.crt", "server.key")
	if err != nil {
		log.Fatalf("failed to load TLS certificates: %v", err)
	}

	list, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer(grpc.Creds(creds))

	proto.RegisterCustomerServiceServer(server, customer.NewCustomerService([]customer.Customer{}))

	fmt.Println("Server running on port :50051")

	log.Printf("server listening at %v", list.Addr())
	if err := server.Serve(list); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
