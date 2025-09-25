// Per default i canali sono _non bufferizzati_, il che significa che
// accetteranno invii (`chan <-`) solo se c'è una
// ricezione corrispondente (`<- chan`) pronta a ricevere il
// valore inviato. I _canali bufferizzati_ accettano un numero
// limitato di valori senza un ricevente corrispondente per
// quei valori.

package main

import "fmt"

func main() {

	// Qui creiamo con `make` un canale di stringhe bufferizzando fino a
	// 2 valori.
	messages := make(chan string, 2)

	// Poiché questo canale è bufferizzato, possiamo inviare questi
	// valori nel canale senza una ricezione
	// concorrente corrispondente.
	messages <- "buffered"
	messages <- "channel"

	// Successivamente possiamo ricevere questi due valori come al solito.
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
