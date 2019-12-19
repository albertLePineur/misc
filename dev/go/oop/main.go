package main

import "fmt"

func main() {
	var monCalcul Calcul = Rectangle{10.5, 7.3}
	fmt.Println(monCalcul.Aire())
	fmt.Println(monCalcul.Perimetre())

	// Initialisation d'une variable de type entier
	var a int = 10
	// Initialisation d'un pointeur de type entier
	var p *int
	// On assigne au pointeur l'adresse mémoire de la variable 'a'
	p = &a
	// equivalent court
	// a := 10
	// p1 := &a

	var pp = &p

	fmt.Println("a = ", a)
	fmt.Println("&a = ", &a)
	fmt.Println("p = ", p)
	fmt.Println("&p = ", &p)
	fmt.Println("*p = ", *p)
	fmt.Println("pp = ", pp)
	fmt.Println("&pp = ", &pp)
	fmt.Println("*pp = ", *pp)
	fmt.Println("**pp = ", **pp)
	*p = 11
	**pp = 12
	fmt.Println("a = ", a)
	fmt.Println("&a = ", &a)
	fmt.Println("p = ", p)
	fmt.Println("&p = ", &p)
	fmt.Println("*p = ", *p)
	fmt.Println("pp = ", pp)
	fmt.Println("&pp = ", &pp)
	fmt.Println("*pp = ", *pp)
	fmt.Println("**pp = ", **pp)

	var prout = Test{2, 3}
	fmt.Println(prout.Method1())

	var monAlbum = Album{
		artiste:        "AC/DC",
		label:          "Albert / Atlantic",
		nomAlbum:       "Back In Black",
		anneeDeSortie:  1980,
		parutionCD:     false,
		parutionVinyle: true,
		lineup: Membres{
			singer:      "Brian Johnson",
			leadGuitar:  "Angus Young",
			rythmGuitar: "Malcolm Young",
			bass:        "Cliff Williams",
			drums:       "Phil Rudd",
		},
	}
	fmt.Println(monAlbum.lineup)
}
