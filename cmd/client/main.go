package main

import (
	"context"
	"log"

	"github.com/jimmysfernandes/grpc-api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("server.crt", "")
	if err != nil {
		log.Fatalf("failed to load TLS certificates: %v", err)
	}

	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := proto.NewCustomerServiceClient(conn)

	customerRequest := &proto.CustomerRequest{
		FirstName: "John",
		LastName:  "Doe",
		Gender:    "M",
		Email:     "jd@email.com",
	}

	resp, err := client.Create(context.Background(), customerRequest)
	if err != nil {
		panic(err)
	}

	log.Printf("Customer created with id: %s", resp.Id)

	client.FindById(context.Background(), &proto.CustomerId{Id: resp.Id})

	log.Printf("Customer found with id: %s", resp.Id)
}
