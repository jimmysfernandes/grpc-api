package customer

import (
	"context"
	"fmt"
	"log"

	"github.com/jimmysfernandes/grpc-api/proto"
)

type CustomerService struct {
	proto.UnimplementedCustomerServiceServer
	customers []Customer
}

func (c *CustomerService) Create(ctx context.Context, req *proto.CustomerRequest) (*proto.CustomerResponse, error) {
	log.Default().Printf("Creating customer with email %s", req.Email)

	customer := NewCustomer(req.FirstName, req.LastName, req.Email)

	c.customers = append(c.customers, *customer)

	log.Default().Printf("Customer created with id %s", customer.Id)

	return &proto.CustomerResponse{
		Id:        customer.Id,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
		Enabled:   customer.Enabled(),
	}, nil
}

func (c *CustomerService) FindById(ctx context.Context, req *proto.CustomerId) (*proto.CustomerResponse, error) {
	log.Default().Printf("Finding customer with id %s", req.Id)

	for _, customer := range c.customers {
		if customer.Id == req.Id {
			log.Default().Printf("Customer with id %s found", req.Id)

			return &proto.CustomerResponse{
				Id:        customer.Id,
				FirstName: customer.FirstName,
				LastName:  customer.LastName,
				Email:     customer.Email,
				Enabled:   customer.Enabled(),
			}, nil
		}
	}

	log.Default().Printf("Customer with id %s not found", req.Id)

	return nil, fmt.Errorf("customer with id %s not found", req.Id)
}

func NewCustomerService(customers []Customer) proto.CustomerServiceServer {
	return &CustomerService{
		customers: customers,
	}
}
