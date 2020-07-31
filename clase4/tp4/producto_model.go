package main

// Producto contiene el nombre y el monto para un super
type Producto struct {
	Tienda string `json:"tienda"`
	ID     int    `json:"id"`
	Precio int    `json:"precio"`
}

// GetPrecio completa la info del producto
func (p *Producto) GetPrecio() error {
	return APIGetProducto(p)
}
