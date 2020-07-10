package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println()
	fmt.Println("Referencia")
	fmt.Println("-----------")

	a := 5
	b := &a
	*b += 2

	fmt.Println(a, *b, b)
	var c *int
	fmt.Println(c)

	fmt.Println()
	fmt.Println("Array")
	fmt.Println("-----------")

	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	fmt.Println()
	fmt.Println("Slice")
	fmt.Println("-----------")

	var s = []int{1, 2, 3}
	fmt.Println(s)
	s[1] = 55
	s = append(s, 88)
	fmt.Println(s)
	s2 := arr[1:]
	s3 := arr[1:]
	s4 := append(s3, 8)
	s3[0] = 100
	fmt.Println(s2, s3, s4)
	fmt.Println(len(s4), cap(s4))
	s5 := arr[1:]
	fmt.Println(s5, len(s5), cap(s5))

	fmt.Println()
	fmt.Println("Map")
	fmt.Println("-----------")

	m1 := map[string]string{"foo": "bar"}
	m1["hola"] = "mundo"
	fmt.Println(m1)
	delete(m1, "hola")
	fmt.Println(m1)

	m1["hola"] = "mundo"
	v, found := m1["hola"]
	fmt.Println(found, v, m1)
	delete(m1, "hola")
	v, found = m1["hola"]
	fmt.Println(found, v, m1)
	m1["pepe"] = "HI!"
	fmt.Println(m1)

	sint := []int{1, 2, 3, 4, 5, 6}
	mss := map[string]string{"a": "a", "b": "b", "c": "c"}
	fmt.Println(sint, mss)
	for i, v := range sint[:2] {
		fmt.Println(i, v)
	}
	for i, v := range mss {
		fmt.Println(i, v)
	}

	fmt.Println()
	fmt.Println("Struct")
	fmt.Println("-----------")

	cir := Circulo{Radio: 1.2, X: 100, Y: 200}
	fmt.Println(cir)
	fmt.Printf("%#v\n", cir)
	fmt.Println(cir.Radio)
	fmt.Println()

	cir2 := NewCirculo()
	cir2.SetRadio(1.5)
	fmt.Println(cir2.Area())
	fmt.Println(cir2)
	ImprimirArea(cir2)

	rect := &Rectangulo{Ancho: 1.3, Alto: 1.5}
	ImprimirArea(rect)
}

// Circulo struct
type Circulo struct {
	Radio float64
	X, Y  int
}

// Area returns de circle area
func (c *Circulo) Area() float64 {
	return math.Pi * math.Pow(c.Radio, 2)
}

// SetRadio sets Circle radio
func (c *Circulo) SetRadio(newRadio float64) {
	c.Radio = newRadio
}

// NewCirculo returns a reference to a new circulo
func NewCirculo() *Circulo {
	return &Circulo{Radio: 1, X: 1, Y: 1}
}

// Rectangulo struct
type Rectangulo struct {
	Ancho, Alto float64
}

// Area returns de Rectangulo area
func (r *Rectangulo) Area() float64 {
	return r.Alto * r.Ancho
}

// Figura interface
type Figura interface {
	Area() float64
}

// ImprimirArea imprime el area de una Figura
func ImprimirArea(f Figura) {
	fmt.Printf("El area es: %f\n", f.Area())
}
