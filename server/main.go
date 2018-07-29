package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	server, err := NewServer()
	if err != nil {
		log.Fatalln(err)
		return
	}
	server.Route(router)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("PORT is invalid")
		return
	}

	log.Println("Listening on http://localhost:" + port)
	log.Fatalln(http.ListenAndServe(":"+port, router))
}
