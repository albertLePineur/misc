package main

// Aire method, calculate the area and returns a float64
func (r Rectangle) Aire() float64 {
	return r.largeur * r.longueur
}

// Perimetre method, calculate the perimeter and returns a float64
func (r Rectangle) Perimetre() float64 {
	return 2 * (r.largeur + r.longueur)
}

// Method1 method
func (t *Test) Method1() int {
	t.champ1 = 10
	return t.champ1 + t.champ2
}
