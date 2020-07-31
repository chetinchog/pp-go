package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const baseURL = "https://productos-p6pdsjmljq-uc.a.run.app"

// APIGetProducto busca en API la info del producto
func APIGetProducto(p *Producto) error {
	url := fmt.Sprintf("%s/%s/productos/%d", baseURL, p.Tienda, p.ID)
	res, errRes := http.Get(url)
	if errRes != nil {
		fmt.Println("Error APIGetProducto:", errRes)
		return errRes
	}
	if res.StatusCode != http.StatusOK {
		fmt.Println("Error APIGetProducto:", res.StatusCode)
		return errRes
	}
	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(p)
	return nil
}
