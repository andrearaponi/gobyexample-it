// Le _interfacce_ sono collezioni nominate di
// signature di metodi.

package main

import (
	"fmt"
	"math"
)

// Ecco un'interfaccia di base per forme geometriche.
type geometry interface {
	area() float64
	perim() float64
}

// Per il nostro esempio implementeremo questa interfaccia sui
// tipi `rect` e `circle`.
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// Per implementare un'interfaccia in Go, dobbiamo solo
// implementare tutti i metodi nell'interfaccia. Qui
// implementiamo `geometry` sui `rect`.
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// L'implementazione per i `circle`.
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// Se una variabile ha un tipo interfaccia, allora possiamo chiamare
// metodi che sono nell'interfaccia nominata. Ecco una
// funzione generica `measure` che sfrutta questo
// per lavorare con qualsiasi `geometry`.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// A volte è utile conoscere il tipo runtime di un
// valore interfaccia. Un'opzione è usare una *type assertion*
// come mostrato qui; un'altra è un [type `switch`](switch).
func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	}
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// I tipi struct `circle` e `rect` entrambi
	// implementano l'interfaccia `geometry` quindi possiamo usare
	// istanze di
	// queste struct come argomenti per `measure`.
	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}
