package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// MethodGet fonction pour générer un résultat lors d'un requête GET
func MethodGet(ecriture http.ResponseWriter, requete *http.Request) {
	valeur := RedisGetKey("macle")

	ecriture.Header().Set("Content-Type", "application/json")
	ecriture.WriteHeader(http.StatusOK)
	ecriture.Write([]byte(valeur))
}

// MethodPost fonction pour générer un résultat lors d'un requête POST
func MethodPost(ecriture http.ResponseWriter, requete *http.Request) {
	var dataByte []byte
	var erreur error
	dataByte, erreur = ioutil.ReadAll(requete.Body)
	if erreur != nil {
		panic(erreur)
	}
	var brut map[string]interface{}
	erreur = json.Unmarshal(dataByte, &brut)
	if erreur != nil {
		panic(erreur)
	}
	fmt.Println(brut)

	ecriture.Header().Set("Content-Type", "application/json")
	ecriture.WriteHeader(http.StatusCreated)
	erreur = RedisSetKey("prout", "caca")
	if erreur != nil {
		ecriture.Write([]byte("Error while writting data to Redis Database"))
	} else {
		ecriture.Write([]byte("Data written to Redis Database"))
	}
}

// MethodPut fonction pour générer un résultat lors d'un requête PUT
func MethodPut(ecriture http.ResponseWriter, requete *http.Request) {
	ecriture.Header().Set("Content-Type", "application/json")
	ecriture.WriteHeader(http.StatusAccepted)
	ecriture.Write([]byte(`{"message": "HTTP PUT"}`))
}

// MethodDelete fonction pour générer un résultat lors d'un requête DELETE
func MethodDelete(ecriture http.ResponseWriter, requete *http.Request) {
	ecriture.Header().Set("Content-Type", "application/json")
	ecriture.WriteHeader(http.StatusOK)
	ecriture.Write([]byte(`{"message": "HTTP DELETE"}`))
}

// MethodNotFound fonction pour générer un résultat lors d'un requête en erreur
func MethodNotFound(ecriture http.ResponseWriter, requete *http.Request) {
	ecriture.Header().Set("Content-Type", "application/json")
	ecriture.WriteHeader(http.StatusNotFound)
	ecriture.Write([]byte(`{"message": "Not Found"}`))
}
