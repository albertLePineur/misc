package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var routeur = mux.NewRouter()
	var apiv1 = routeur.PathPrefix("/api/v1").Subrouter()
	apiv1.HandleFunc("", MethodGet).Methods(http.MethodGet)
	apiv1.HandleFunc("", MethodPost).Methods(http.MethodPost)
	apiv1.HandleFunc("", MethodPut).Methods(http.MethodPut)
	apiv1.HandleFunc("", MethodDelete).Methods(http.MethodDelete)
	apiv1.HandleFunc("", MethodNotFound)
	//apiv1.HandleFunc("/user/{userID}/comment/{commentID}", parametres).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", routeur))
}
