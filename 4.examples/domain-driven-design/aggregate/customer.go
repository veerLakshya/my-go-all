// this package contains aggregates which combine multiple entities into single object
package aggregate

import (
	"errors"

	"github.com/google/uuid"
	"github.com/veerlakshya/my-go-all/4.examples/ddd/entity"
	"github.com/veerlakshya/my-go-all/4.examples/ddd/valueObject"
)

var (
	ErrInvalidPerson = errors.New("a customer must have valid name")
)

type Customer struct {
	// Person is the root identity of this object
	// which means Person.Id is the main identifier for the customer
	person       *entity.Person
	products     []*entity.Item
	transactions []valueObject.Transaction
}

func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	newPerson := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}

	return Customer{
		person:       newPerson,
		products:     make([]*entity.Item, 0),
		transactions: make([]valueObject.Transaction, 0),
	}, nil
}

func (c *Customer) GetId() uuid.UUID {
	return c.person.ID
}

func (c *Customer) GetName() string {
	return c.person.Name
}

func (c *Customer) SetId(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.ID = id
}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.Name = name
}
