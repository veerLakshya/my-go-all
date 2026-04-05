// this package holds all entities which are shared across subdomains
package entity

import (
	"github.com/google/uuid"
)

// represents person in all domains
type Person struct {
	ID   uuid.UUID
	Name string
	Age  int
}
