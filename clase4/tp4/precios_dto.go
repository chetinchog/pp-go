package main

// PreciosDTOR (Read) estructura de entrada
type PreciosDTOR struct {
	ListID    []int    `json:"list_id"`
	ListSuper []string `json:"list_super"`
}

// PreciosDTOW (Write) contiene el nombre y el monto para un super
type PreciosDTOW struct {
	Tienda string `json:"tienda"`
	Precio int    `json:"precio"`
}
