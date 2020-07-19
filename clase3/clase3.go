package main

import (
	"fmt"
	"time"
)

type ae struct {
}

func (i ae) Error() string {
	return "Error ae"
}

type be struct {
}

func (i be) Error() string {
	return "Error be"
}

// Sumar suma :D
func Sumar(a, b int) (int, error) {
	return 0, ae{}
	// return 0, fmt.Errorf("No pude sumar %d + %d", a, b)
	// return a + b, nil
}

func showError() {
	n, err := Sumar(1, 2)
	if _, ok := err.(ae); ok {
		fmt.Println("Error a")
	}
	if _, ok := err.(be); ok {
		fmt.Println("Error b")
	}
	// if err != nil {
	// 	log.Printf("Hubo un porblema al sumar: %v\n", err)
	// 	// os.Exit(-1)
	// }
	fmt.Println(n)
}

func imprimir(s string) {
	defer fmt.Printf("Imprimir defer %s\n", s)
	fmt.Printf("Imprimir %s\n", s)
}

func deferers() {
	defer imprimir("1")
	defer imprimir("2")
	fmt.Println("Estoy haciendo algo")
	defer imprimir("3")
	defer imprimir("4")
	defer func(s string) {
		fmt.Println("Funcion anonima", s)
	}("A")
}

func impGoRoutine() {
	fmt.Println("impGoRoutine!")
}

func do() {
	time.Sleep(time.Second * 1)
	fmt.Println("DO")
}

func goChan(foo func(), done chan bool) {
	go func() {
		foo()
		done <- true
	}()
}

func goAwait(foo func()) {
	done := make(chan bool)
	goChan(foo, done)
	<-done
}

func concurrencia() {
	fmt.Println("Concurrencia")
	go impGoRoutine()
	fmt.Println("Foo!")
	time.Sleep(time.Second)
	fmt.Println()

	msgsChan := make(chan string)
	// Read: s := <- msgsChan
	// Write: msgsChan <- "Hi!"
	fmt.Println("Hola")
	go func() {
		time.Sleep(time.Second * 1)
		msgsChan <- "Mundo"
	}()
	s := <-msgsChan
	fmt.Println(s)
	fmt.Println()

	messages := make(chan string, 2)

	messages <- "Hola"
	messages <- "Mundo"
	out := <-messages
	messages <- "!"
	fmt.Println(out)
	fmt.Println()

	done := make(chan bool)
	fmt.Println("GO!")
	goChan(do, done)
	fmt.Println("Wait...")
	<-done
	fmt.Println("Done")
	goAwait(do)
	fmt.Println("Done2")
}

func main() {
	fmt.Println("Clase3")
	showError()
	fmt.Println()
	deferers()
	fmt.Println()
	concurrencia()
}
