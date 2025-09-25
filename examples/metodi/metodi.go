// Go supporta i _metodi_ definiti sui tipi struct.

package main

import "fmt"

type rect struct {
	width, height int
}

// Questo metodo `area` ha un _tipo receiver_ di `*rect`.
func (r *rect) area() int {
	return r.width * r.height
}

// I metodi possono essere definiti sia per puntatori che per tipi
// receiver di valore. Ecco un esempio di receiver di valore.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// Qui chiamiamo i 2 metodi definiti per la nostra struct.
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Go gestisce automaticamente la conversione tra valori
	// e puntatori per le chiamate di metodo. Potresti voler usare
	// un tipo receiver puntatore per evitare la copia nelle chiamate
	// di metodo o per permettere al metodo di mutare la
	// struct ricevente.
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
