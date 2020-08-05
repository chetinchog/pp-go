package product

import (
	"fmt"

	"github.com/google/uuid"
)

func init() {
	fmt.Println("> Product init")
}

// Product struct
type Product struct {
	ID    string
	Name  string
	Price float32
}

// NewProduct returns new Product
func NewProduct(name string, price float32) (*Product, error) {
	if price <= 0 {
		return nil, fmt.Errorf(
			"No se puede crear Product con price menor a 0. Recibido con %f",
			price,
		)
	}
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}, nil
}
