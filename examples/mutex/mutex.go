// Nell'esempio precedente abbiamo visto come gestire semplici
// stati di contatori usando [operazioni atomiche](contatori-atomici).
// Per stati più complessi possiamo usare un [_mutex_](https://en.wikipedia.org/wiki/Mutual_exclusion)
// per accedere ai dati in modo sicuro attraverso più goroutine.

package main

import (
	"fmt"
	"sync"
)

// Container contiene una mappa di contatori; dato che vogliamo
// aggiornarla simultaneamente da più goroutine, aggiungiamo
// un `Mutex` per sincronizzare l'accesso.
// Nota che i mutex non devono essere copiati, quindi se questa
// `struct` viene passata in giro, dovrebbe essere fatto tramite
// puntatore.
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	// Blocchiamo il mutex prima di accedere a `counters`; lo sblocchiamo
	// alla fine della funzione usando un'istruzione [defer](defer).
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	c := Container{
		// Nota che il valore zero di un mutex è utilizzabile così com'è, quindi
		// non è richiesta alcuna inizializzazione qui.
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	// Questa funzione incrementa un contatore nominato
	// in un ciclo.
	doIncrement := func(name string, n int) {
		for range n {
			c.inc(name)
		}
	}

	// Eseguiamo diverse goroutine simultaneamente; nota
	// che tutte accedono allo stesso `Container`,
	// e due di esse accedono allo stesso contatore.
	wg.Go(func() {
		doIncrement("a", 10000)
	})

	wg.Go(func() {
		doIncrement("a", 10000)
	})

	wg.Go(func() {
		doIncrement("b", 10000)
	})

	// Aspettiamo che le goroutine finiscano
	wg.Wait()
	fmt.Println(c.counters)
}
