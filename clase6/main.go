package main

import (
	"fmt"

	"github.com/chetinchog/pp-go/clase6/product"
)

func prod() {
	var pro product.Product = product.Product{
		ID:    "1",
		Name:  "Alcohol",
		Price: 100,
	}
	fmt.Printf("%#v\n", pro)
}

func main() {
	fmt.Println()
	fmt.Println("Clase 6")
	fmt.Println("--------------")
	fmt.Println()

	fmt.Println("Product")
	fmt.Println("--------------")
	prod()
	fmt.Println()
}
