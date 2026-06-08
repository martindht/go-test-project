package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr  string
	store Store
}

// constructor for server
func NewAPIServer(addr string, store Store) *APIServer {
	return &APIServer{addr: addr, store: store}
}

// method server has. initialize router
func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// register service
	tasksService := NewTasksService(s.store)
	tasksService.RegisterRoutes(router)
	log.Println("Staring the API server at", s.addr)

	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}
