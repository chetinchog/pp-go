package main

import (
	"fmt"
)

func sumar(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Hola mundo!")
	s := sumar
	r := s(10, 5)
	fmt.Println(r)

	// var a int = 10
	// b := 20
	const c = "hola"

	fmt.Println()

	for i := 0; i < 3; i++ {
		fmt.Printf("Hola %d\n", i)
	}

	fmt.Println()

	if m := 2 % 2; m == 0 {
		fmt.Println("Par")
	}
}
