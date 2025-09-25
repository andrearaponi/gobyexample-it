// Il meccanismo primario per gestire lo stato in Go è
// la comunicazione attraverso i canali. L'abbiamo visto per esempio
// con i [pool di worker](pool-di-worker). Ci sono però alcune altre
// opzioni per gestire lo stato. Qui vedremo
// l'uso del pacchetto `sync/atomic` per _contatori atomici_
// accessibili da più goroutine.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	// Useremo un tipo integer atomico per rappresentare il nostro
	// contatore (sempre positivo).
	var ops atomic.Uint64

	// Un WaitGroup ci aiuterà ad attendere che tutte le goroutine
	// finiscano il loro lavoro.
	var wg sync.WaitGroup

	// Avvieremo 50 goroutine che ognuna incrementa il
	// contatore esattamente 1000 volte.
	for range 50 {
		wg.Go(func() {
			for range 1000 {
				// Per incrementare atomicamente il contatore usiamo `Add`.
				ops.Add(1)
			}
		})
	}

	// Aspettiamo fino a quando tutte le goroutine hanno finito.
	wg.Wait()

	// Qui nessuna goroutine sta scrivendo su 'ops', ma usando
	// `Load` è sicuro leggere atomicamente un valore anche mentre
	// altre goroutine lo stanno (atomicamente) aggiornando.
	fmt.Println("ops:", ops.Load())
}
