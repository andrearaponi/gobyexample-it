// Una _goroutine_ è un thread di esecuzione leggero.

package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// Supponiamo di avere una chiamata di funzione `f(s)`. Ecco come
	// la chiameremmo nel modo normale, eseguendola
	// in modo sincrono.
	f("direct")

	// Per invocare questa funzione in una goroutine, usa
	// `go f(s)`. Questa nuova goroutine eseguirà
	// concorrentemente con quella chiamante.
	go f("goroutine")

	// Puoi anche avviare una goroutine per una chiamata
	// di funzione anonima.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Le nostre due chiamate di funzione ora stanno eseguendo
	// asincronamente in goroutine separate. Aspettiamo che finiscano
	// (per un approccio più robusto, usa un [WaitGroup](waitgroups)).
	time.Sleep(time.Second)
	fmt.Println("done")
}
