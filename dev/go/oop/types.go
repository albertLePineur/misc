package main

// Calcul interface
type Calcul interface {
	Aire() float64      // Call to Aire method (it takes no arguments and returns a float64)
	Perimetre() float64 // Call to Perimetre method (it takes no arguments and returns a float64)
}

// Rectangle structure
type Rectangle struct {
	largeur  float64
	longueur float64
}

// Test type
type Test struct {
	champ1 int
	champ2 int
}

// Album structure
type Album struct {
	artiste, label, nomAlbum   string
	anneeDeSortie              int16
	parutionCD, parutionVinyle bool
	lineup                     Membres
}

// Membres structure
type Membres struct {
	singer, leadGuitar, rythmGuitar, bass, drums string
}
