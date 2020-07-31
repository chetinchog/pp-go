package main

import (
	"encoding/json"
	"net/http"
)

// PreciosView recibe una lista de IDs de productos y una lista de supermercados.
// La respuesta tiene que ser la suma de precios para cada uno de los supermercados.
func PreciosView(w http.ResponseWriter, r *http.Request) {
	// Asegura cerrar Body
	defer r.Body.Close()
	// Convierte JSON al DTO correspondiente
	var consultaDTO PreciosDTOR
	json.NewDecoder(r.Body).Decode(&consultaDTO)
	// Ejecuta controlador
	listPrecios, err := PreciosController(consultaDTO.ListID, consultaDTO.ListSuper)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	// Convierte la respuesta al DTO correspondiente
	var respuestaDTO []PreciosDTOW
	for _, precio := range listPrecios {
		respuestaDTO = append(respuestaDTO, PreciosDTOW{
			Precio: precio.Precio,
			Tienda: precio.Tienda,
		})
	}
	// Responde la peticion
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respuestaDTO)
}
