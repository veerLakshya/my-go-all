package aggregate_test

import (
	"errors"
	"testing"

	"github.com/veerlakshya/my-go-all/4.examples/ddd/aggregate"
)

func TestCustomer_NewCustomer(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Empty name test",
			name:        "",
			expectedErr: aggregate.ErrInvalidPerson,
		},
		{
			test:        "valid name test",
			name:        "lakshya",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewCustomer(tc.name)

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error: %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
