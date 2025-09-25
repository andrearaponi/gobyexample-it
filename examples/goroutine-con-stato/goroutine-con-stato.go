// Nell'esempio precedente abbiamo usato il locking esplicito con
// [mutex](mutex) per sincronizzare l'accesso allo stato condiviso
// attraverso più goroutine. Un'altra opzione è utilizzare le
// funzionalità di sincronizzazione integrate delle goroutine e
// dei canali per ottenere lo stesso risultato. Questo approccio basato
// sui canali si allinea con le idee di Go di condividere la memoria
// comunicando e avere ogni pezzo di dati posseduto
// esattamente da 1 goroutine.

package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// In questo esempio il nostro stato sarà posseduto da una singola
// goroutine. Questo garantirà che i dati non vengano mai
// corrotti dall'accesso simultaneo. Per leggere o
// scrivere tale stato, altre goroutine invieranno messaggi
// alla goroutine proprietaria e riceveranno le corrispondenti
// risposte. Queste `struct` `readOp` e `writeOp`
// incapsulano quelle richieste e un modo per la goroutine
// proprietaria di rispondere.
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	// Come prima conteremo quante operazioni eseguiamo.
	var readOps uint64
	var writeOps uint64

	// I canali `reads` e `writes` saranno utilizzati da
	// altre goroutine per emettere richieste di lettura e scrittura,
	// rispettivamente.
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// Ecco la goroutine che possiede lo `state`, che
	// è una mappa come nell'esempio precedente ma ora privata
	// alla goroutine con stato. Questa goroutine ripetutamente
	// fa select sui canali `reads` e `writes`,
	// rispondendo alle richieste man mano che arrivano. Una risposta
	// viene eseguita prima eseguendo l'operazione richiesta
	// e poi inviando un valore sul canale di risposta
	// `resp` per indicare il successo (e il valore desiderato
	// nel caso di `reads`).
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// Questo avvia 100 goroutine per emettere letture alla
	// goroutine proprietaria dello stato tramite il canale `reads`.
	// Ogni lettura richiede la costruzione di un `readOp`, inviandolo
	// attraverso il canale `reads`, e poi ricevendo il
	// risultato attraverso il canale `resp` fornito.
	for range 100 {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Avviamo anche 10 scritture, usando un approccio
	// simile.
	for range 10 {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Lasciamo lavorare le goroutine per un secondo.
	time.Sleep(time.Second)

	// Infine, catturiamo e riportiamo i conteggi delle operazioni.
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
