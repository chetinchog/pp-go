package main

import (
	"fmt"

	"github.com/chetinchog/pp-go/clase6/product"
	"github.com/fatih/color"
)

func prod() {
	var pro product.Product = product.Product{
		ID:    "1",
		Name:  "Alcohol",
		Price: 100,
	}
	fmt.Printf("%#v\n", pro)
	color.Red(fmt.Sprintf("%#v", pro))
	color.Green(fmt.Sprintf("%#v", pro))
	color.Blue(fmt.Sprintf("%#v", pro))
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
