package product

import "fmt"

func init() {
	fmt.Println("> Product init")
}

// Product struct
type Product struct {
	ID    string
	Name  string
	Price float32
}
