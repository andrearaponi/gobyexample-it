// Quando usi i canali come parametri di funzione, puoi
// specificare se un canale è destinato solo a inviare o ricevere
// valori. Questa specificità aumenta la type-safety del
// programma.

package main

import "fmt"

// Questa funzione `ping` accetta solo un canale per inviare
// valori. Sarebbe un errore a compile-time provare a
// ricevere su questo canale.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// La funzione `pong` accetta un canale per le ricezioni
// (`pings`) e un secondo per gli invii (`pongs`).
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
