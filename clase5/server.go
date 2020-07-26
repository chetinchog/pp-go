package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type handler struct{}

func (h handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	if req.URL.Path != "/foo" {
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	defer req.Body.Close()
	var m map[string]string
	json.NewDecoder(req.Body).Decode(&m)
	fmt.Printf("%#v\n", m)

	lala := map[string]string{
		"foo": "bar",
	}
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(lala)
}

// Server run
func Server() {
	http.ListenAndServe(":8080", handler{})
}
