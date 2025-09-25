// Possiamo usare i canali per sincronizzare l'esecuzione
// tra goroutine. Ecco un esempio di utilizzo di una
// ricezione bloccante per aspettare che una goroutine finisca.
// Quando si aspettano più goroutine per finire,
// potresti preferire usare un [WaitGroup](waitgroups).

package main

import (
	"fmt"
	"time"
)

// Questa è la funzione che eseguiremo in una goroutine. Il
// canale `done` sarà usato per notificare a un'altra
// goroutine che il lavoro di questa funzione è terminato.
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// Invia un valore per notificare che abbiamo finito.
	done <- true
}

func main() {

	// Avvia una goroutine worker, passandole il canale su cui
	// notificare.
	done := make(chan bool, 1)
	go worker(done)

	// Blocca finché riceviamo una notifica dal
	// worker sul canale.
	<-done
}
