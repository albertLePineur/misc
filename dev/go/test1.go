package main

import (
	"fmt"
	"time"
)

// MaFonction function
func MaFonction(nom string, entier int, temps time.Duration, canal chan int) {
	for i := 0; i < 10; i++ {
		entier = (entier + (entier - 1))
		fmt.Println(nom, entier)
		time.Sleep(temps * time.Millisecond)
	}
	canal <- entier
}

func main() {

	// Initialisation d'un cana d'entier monCanal
	var monCanal = make(chan int)

	// On execute deux fois la même fonction en mode Goroutine (process indépendant)
	go MaFonction("test1", 2, 2000, monCanal)
	go MaFonction("test2", 3, 1500, monCanal)

	// On vient verser le contenu du channel monCanal dans les deux variables ci-dessous
	var resultat1 = <-monCanal
	var resultat2 = <-monCanal

	// On vient afficher le contenu des canaux
	fmt.Println(resultat1, resultat2)
}
