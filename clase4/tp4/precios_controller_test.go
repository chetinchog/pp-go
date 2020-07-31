package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrecios_CalcularPrecios(t *testing.T) {
	casos := []struct {
		nombre    string
		listID    []int
		listSuper []string
		esperado  []Carrito
	}{
		{
			nombre:    "Calcular el precio correcto por supermercado",
			listID:    []int{0, 1, 2, 3, 4, 5},
			listSuper: []string{"Jumbo", "SuperVea", "Carrefour", "Macro"},
			esperado: []Carrito{
				Carrito{Tienda: "Jumbo", Precio: 38946},
				Carrito{Tienda: "SuperVea", Precio: 31687},
				Carrito{Tienda: "Carrefour", Precio: 24771},
				Carrito{Tienda: "Macro", Precio: 50675},
			},
		},
		{
			nombre:    "Retornar vacio cuando no hay productos ni supers",
			listID:    []int{},
			listSuper: []string{},
			esperado:  []Carrito{},
		},
	}

	for _, test := range casos {
		t.Run(test.nombre, func(t *testing.T) {
			respuesta, err := PreciosController(test.listID, test.listSuper)
			if err != nil {
				fmt.Printf("Error: %#v\n", err)
				return
			}
			if len(respuesta) != len(test.esperado) {
				t.Errorf("CalcularPrecios retorno %d supermercados, se esperaban %d", len(respuesta), len(test.esperado))
			}

			if !reflect.DeepEqual(respuesta, test.esperado) {
				t.Errorf("CalcularPrecios retorna precios incorrectos %+v, se esperaban %+v", respuesta, test.esperado)
			}
		})
	}
}
