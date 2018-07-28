package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	server := NewServer()
	server.Route(router)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("PORT is invalid")
	}

	log.Println("Listening on http://localhost:" + port)
	log.Fatalln(http.ListenAndServe(":"+port, router))
}
