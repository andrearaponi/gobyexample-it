// Gli invii e le ricezioni di base sui canali sono bloccanti.
// Tuttavia, possiamo usare `select` con una clausola `default` per
// implementare invii _non-bloccanti_, ricezioni, e persino
// `select` multi-way non-bloccanti.

package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// Ecco una ricezione non-bloccante. Se un valore è
	// disponibile su `messages` allora `select` prenderà
	// il `case` `<-messages` con quel valore. Se no
	// prenderà immediatamente il `case` `default`.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// Un invio non-bloccante funziona similarmente. Qui `msg`
	// non può essere inviato al canale `messages`, perché
	// il canale non ha buffer e non c'è nessun ricevente.
	// Quindi viene selezionato il `case` `default`.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// Possiamo usare multipli `case` sopra la clausola `default`
	// per implementare un select multi-way non-bloccante.
	// Qui proviamo ricezioni non-bloccanti sia su
	// `messages` che su `signals`.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
