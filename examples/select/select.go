// Il _select_ di Go ti permette di aspettare su multiple
// operazioni di canale. Combinare goroutine e canali con
// select è una funzionalità potente di Go.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Per il nostro esempio faremo select su due canali.
	c1 := make(chan string)
	c2 := make(chan string)

	// Ogni canale riceverà un valore dopo un certo
	// tempo, per simulare ad es. operazioni RPC bloccanti
	// in esecuzione in goroutine concorrenti.
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// Useremo `select` per aspettare entrambi questi valori
	// simultaneamente, stampando ognuno non appena arriva.
	for range 2 {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
