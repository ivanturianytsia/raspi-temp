package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	server := NewServer()
	server.Route(router)

	http.ListenAndServe(":3000", router)
}
