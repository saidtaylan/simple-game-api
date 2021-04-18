package main

import (
	"github.com/gorilla/mux"
)

func multiplexer() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/games", getGames).Methods("GET")
	r.HandleFunc("game/{id}", getGame).Methods("GET")
	r.HandleFunc("/games", createGame).Methods("POST")
	r.HandleFunc("/games/{id}", updateGame).Methods("PUT")
	r.HandleFunc("/games/{id}", deleteGame).Methods("DELETE")
	return r
}
