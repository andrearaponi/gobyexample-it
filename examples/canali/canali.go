// I _canali_ sono le pipe che collegano le goroutine
// concorrenti. Puoi inviare valori nei canali da una
// goroutine e ricevere quei valori in un'altra
// goroutine.

package main

import "fmt"

func main() {

	// Crea un nuovo canale con `make(chan val-type)`.
	// I canali sono tipizzati dai valori che trasportano.
	messages := make(chan string)

	// _Invia_ un valore in un canale usando la sintassi `channel <-`.
	// Qui inviamo `"ping"` al canale `messages`
	// che abbiamo creato sopra, da una nuova goroutine.
	go func() { messages <- "ping" }()

	// La sintassi `<-channel` _riceve_ un valore dal
	// canale. Qui riceveremo il messaggio `"ping"`
	// che abbiamo inviato sopra e lo stamperemo.
	msg := <-messages
	fmt.Println(msg)
}
