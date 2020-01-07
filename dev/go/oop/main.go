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
	// & est l'adresse de l'opérateur

	fmt.Println("a = ", a)       // a == 10
	fmt.Println("&a = ", &a)     // &a == <adresse_memoire_a> (sous forme 0x00000...)
	fmt.Println("p = ", p)       // p == <adresse_memoire_a>
	fmt.Println("&p = ", &p)     // &p == <adresse_memoire_p>
	fmt.Println("*p = ", *p)     // *p == 10 (le pointeur renvoie vers )
	fmt.Println("pp = ", pp)     // pp == <adresse_memoire_p>
	fmt.Println("&pp = ", &pp)   // &pp == <adresse_memoire_pp>
	fmt.Println("*pp = ", *pp)   // *pp == <adresse_memoire_a>
	fmt.Println("**pp = ", **pp) // **pp == 10

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
