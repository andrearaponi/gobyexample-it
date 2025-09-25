// A volte vorremmo che i nostri programmi Go gestiscano intelligentemente
// i [segnali Unix](https://en.wikipedia.org/wiki/Unix_signal).
// Per esempio, potremmo volere che un server si spenga
// gracefully quando riceve un `SIGTERM`, o che uno strumento
// da riga di comando smetta di processare input se riceve un `SIGINT`.
// Ecco come gestire i segnali in Go con i canali.

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// La notifica dei segnali di Go funziona inviando valori
	// `os.Signal` su un canale. Creeremo un canale per
	// ricevere queste notifiche. Nota che questo canale
	// dovrebbe essere bufferizzato.
	sigs := make(chan os.Signal, 1)

	// `signal.Notify` registra il canale dato per
	// ricevere notifiche dei segnali specificati.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Potremmo ricevere da `sigs` qui nella funzione
	// main, ma vediamo come questo possa essere fatto
	// anche in una goroutine separata, per dimostrare
	// uno scenario più realistico di spegnimento graceful.
	done := make(chan bool, 1)

	go func() {
		// Questa goroutine esegue una ricezione bloccante per
		// i segnali. Quando ne riceve uno lo stamperà
		// e poi notificherà al programma che può finire.
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	// Il programma aspetterà qui fino a quando non riceve il
	// segnale atteso (come indicato dalla goroutine
	// sopra che invia un valore su `done`) e poi uscirà.
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
