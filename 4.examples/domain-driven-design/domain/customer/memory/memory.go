package memory

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/veerlakshya/my-go-all/4.examples/ddd/aggregate"
	"github.com/veerlakshya/my-go-all/4.examples/ddd/domain/customer"
)

type MemoryRepository struct {
	Customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

func New() *MemoryRepository {

	return &MemoryRepository{
		Customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

func (mr *MemoryRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	c, ok := mr.Customers[id]
	if ok {
		return c, nil
	}
	return aggregate.Customer{}, customer.ErrCustomerNotFound
}

func (mr *MemoryRepository) Add(newCustomer aggregate.Customer) error {
	mr.Lock()
	defer mr.Unlock()

	if mr.Customers == nil {
		mr.Customers = make(map[uuid.UUID]aggregate.Customer)
	}

	_, ok := mr.Customers[newCustomer.GetId()]

	if ok {
		return fmt.Errorf("customer already exists")
	}
	mr.Customers[newCustomer.GetId()] = newCustomer
	return nil
}

func (mr *MemoryRepository) Update(newCustomer aggregate.Customer) error {
	mr.Lock()
	defer mr.Unlock()

	_, ok := mr.Customers[newCustomer.GetId()]

	if !ok {
		return fmt.Errorf("customer doesnt exists")
	}
	//overwrite current one
	mr.Customers[newCustomer.GetId()] = newCustomer
	return nil
}
