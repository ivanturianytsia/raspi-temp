package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	store *DataStore
}

func NewServer() (*Server, error) {
	store, err := NewDataStore()
	if err != nil {
		return nil, err
	}
	return &Server{
		store: store,
	}, nil
}

func (server *Server) Route(router *mux.Router) {
	router.Methods("POST").Path("/api/temp").HandlerFunc(server.handlePostTemp)
	router.Methods("GET").Path("/api/temp").HandlerFunc(server.handleGetTemp)
	router.Methods("GET").Handler(http.FileServer(http.Dir("./dist")))
}

func (server *Server) handlePostTemp(w http.ResponseWriter, r *http.Request) {
	var point DataPoint
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&point); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if err := server.store.Save(point); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if err := json.NewEncoder(w).Encode(&point); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}

func (server *Server) handleGetTemp(w http.ResponseWriter, r *http.Request) {
	point, err := server.store.Get(time.Now())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if err := json.NewEncoder(w).Encode(&point); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}
