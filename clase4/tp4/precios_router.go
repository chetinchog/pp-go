package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// PreciosRouter carga los endpoints
func PreciosRouter(r *mux.Router) {
	s := r.PathPrefix("/precios").Subrouter()
	s.HandleFunc("", PreciosView).
		Methods(http.MethodGet)
	s.HandleFunc("/", PreciosView).
		Methods(http.MethodGet)
}
