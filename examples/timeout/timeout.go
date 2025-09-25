// I _timeout_ sono importanti per programmi che si connettono a
// risorse esterne o che altrimenti necessitano di limitare
// il tempo di esecuzione. Implementare timeout in Go è facile ed
// elegante grazie ai canali e `select`.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Per il nostro esempio, supponiamo di eseguire una chiamata
	// esterna che restituisce il suo risultato su un canale `c1`
	// dopo 2s. Nota che il canale è bufferizzato, quindi l'invio
	// nella goroutine è non-bloccante. Questo è un pattern
	// comune per prevenire leak di goroutine nel caso in cui
	// il canale non venga mai letto.
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// Ecco il `select` che implementa un timeout.
	// `res := <-c1` aspetta il risultato e `<-time.After`
	// aspetta che venga inviato un valore dopo il timeout di
	// 1s. Poiché `select` procede con la prima ricezione
	// che è pronta, prenderemo il caso timeout se l'operazione
	// impiega più di 1s consentito.
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// Se permettiamo un timeout più lungo di 3s, allora la ricezione
	// da `c2` avrà successo e stamperemo il risultato.
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
