package main

import (
	"fmt"
	"sync"
)

type jobR struct {
	Tienda   string
	Producto int
}

type jobW struct {
	Tienda string
	Precio int
}

// PreciosController recibe una lista de IDs de productos y una lista de supermercados.
// La respuesta tiene que ser la suma de precios para cada uno de los supermercados.
func PreciosController(listID []int, listSuper []string) ([]Carrito, error) {
	listCarrito := []Carrito{}
	if len(listID) < 1 || len(listSuper) < 0 {
		return listCarrito, nil
	}

	cantJobs := len(listID) * len(listSuper)
	results := make(chan *Carrito, cantJobs)

	wg := sync.WaitGroup{}
	wg.Add(cantJobs)
	for _, super := range listSuper {
		for _, id := range listID {
			go func(super string, id int) {
				defer wg.Done()
				results <- calcularPrecio(super, id)
			}(super, id)
		}
	}
	go func() {
		defer close(results)
		wg.Wait()
	}()

	mapSuper := make(map[string]int, 0)
	for result := range results {
		mapSuper[result.Tienda] += result.Precio
	}
	for _, super := range listSuper {
		listCarrito = append(listCarrito, Carrito{Tienda: super, Precio: mapSuper[super]})
	}
	fmt.Println("CantJobs:", cantJobs, "->", "cantProd:", len(listID), "cantSuper:", len(listSuper))
	return listCarrito, nil
}

func calcularPrecio(tienda string, id int) *Carrito {
	producto := &Producto{ID: id, Tienda: tienda}
	if err := producto.GetPrecio(); err != nil {
		return &Carrito{Tienda: tienda, Precio: 0}
	}
	return &Carrito{Tienda: tienda, Precio: producto.Precio}
}
