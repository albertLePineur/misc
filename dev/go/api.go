package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type evenement struct {
	id          string `json:"id"`
	titre       string `json:"titre"`
	description string `json:"description"`
}

// TousLesEvenements type
type TousLesEvenements []evenement

var evenements = TousLesEvenements{
	{
		id:          "1",
		titre:       "ma premiere page",
		description: "trololol",
	},
}

// Racine fonction
func Racine(affichage http.ResponseWriter, requete *http.Request) {
	fmt.Fprintf(affichage, "Bienvenue tête de cul !")
}

// CreationEvenement fonction
func CreationEvenement(ecriture http.ResponseWriter, requete *http.Request) {
	var nouvelEvenement evenement
	requeteCorps, erreur := ioutil.ReadAll(requete.Body)
	if erreur != nil {
		fmt.Fprintf(ecriture, "Entrer la donnee avec le titre et la description pour mettre à jour")
	}
	json.Unmarshal(requeteCorps, &nouvelEvenement)
	evenements = append(evenements, nouvelEvenement)
	ecriture.WriteHeader(http.StatusCreated)
	json.NewEncoder(ecriture).Encode(nouvelEvenement)
}

func main() {
	var routeur = mux.NewRouter().StrictSlash(true)
	routeur.HandleFunc("/", Racine)
	routeur.HandleFunc("/evenement", CreationEvenement).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", routeur))
}
