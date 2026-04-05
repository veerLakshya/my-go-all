package customer

import (
	"errors"

	"github.com/google/uuid"
	"github.com/veerlakshya/my-go-all/4.examples/ddd/aggregate"
)

var (
	ErrCustomerNotFound       = errors.New("Customer not found in repo")
	ErrFailedToAddCustomer    = errors.New("Failed to add customer")
	ErrFailedToUpdateCustomer = errors.New("Failed to update customer")
)

type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}
