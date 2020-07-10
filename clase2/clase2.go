package main

import "fmt"

func main() {
	fmt.Println()
	fmt.Println("Puntero")
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
	fmt.Println()

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
	fmt.Println()

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
	fmt.Println()

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
	fmt.Println()
}
