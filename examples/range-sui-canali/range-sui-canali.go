// In un esempio [precedente](range-sui-tipi-built-in) abbiamo visto come `for` e
// `range` forniscano iterazione su strutture dati di base.
// Possiamo anche usare questa sintassi per iterare sui
// valori ricevuti da un canale.

package main

import "fmt"

func main() {

	// Itereremo su 2 valori nel canale `queue`.
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// Questo `range` itera su ogni elemento non appena viene
	// ricevuto da `queue`. Poiché abbiamo chiuso il
	// canale sopra, l'iterazione termina dopo aver
	// ricevuto i 2 elementi.
	for elem := range queue {
		fmt.Println(elem)
	}
}
