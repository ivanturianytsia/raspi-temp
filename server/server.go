package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	store DataStore
}

func NewServer() *Server {
	return &Server{
		store: NewDataStore(),
	}
}

func (server *Server) Route(router *mux.Router) {
	router.Methods("POST").Path("/temp").HandlerFunc(server.handlePostTemp)
	router.Methods("GET").Path("/temp").HandlerFunc(server.handleGetTemp)
}

func (server *Server) handlePostTemp(w http.ResponseWriter, r *http.Request) {
	var point DataPoint
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&point); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	server.store.Set(point)

	if err := json.NewEncoder(w).Encode(&point); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}

func (server *Server) handleGetTemp(w http.ResponseWriter, r *http.Request) {
	point := server.store.Get()

	if err := json.NewEncoder(w).Encode(&point); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}
