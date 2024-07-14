package customer

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	Id        string    `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	EnabledAt time.Time `json:"enabledAt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (c *Customer) Enable() {
	c.EnabledAt = time.Now()
}

func (c *Customer) Enabled() bool {
	return c.EnabledAt != time.Time{}
}

func NewCustomer(firstName, lastName, email string) *Customer {
	return &Customer{
		Id:        uuid.NewString(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
