package memory

import (
	"testing"

	"github.com/google/uuid"
	"github.com/veerlakshya/my-go-all/4.examples/ddd/aggregate"
)

func TestMemory_GetCustom(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	cust, err := aggregate.NewCustomer("lakshya")
}
