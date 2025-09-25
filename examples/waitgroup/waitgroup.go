// Per attendere che più goroutine terminino, possiamo
// utilizzare un *wait group*.

package main

import (
	"fmt"
	"sync"
	"time"
)

// Questa è la funzione che eseguiremo in ogni goroutine.
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	// Aspettiamo per simulare un'operazione costosa.
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	// Questo WaitGroup viene utilizzato per attendere che tutte le
	// goroutine avviate qui terminino. Nota: se un WaitGroup viene
	// passato esplicitamente alle funzioni, dovrebbe essere fatto *per puntatore*.
	var wg sync.WaitGroup

	// Avvia diverse goroutine utilizzando `WaitGroup.Go`
	for i := 1; i <= 5; i++ {
		wg.Go(func() {
			worker(i)
		})
	}

	// Blocca fino a quando tutte le goroutine avviate da `wg` sono
	// terminate. Una goroutine è terminata quando la funzione che invoca
	// ritorna.
	wg.Wait()

	// Nota che questo approccio non ha un modo diretto
	// per propagare errori dai worker. Per casi d'uso più
	// avanzati, considera l'utilizzo del
	// [pacchetto errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup).
}
