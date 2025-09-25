// I [timer](timer) sono per quando vuoi fare
// qualcosa una volta nel futuro - i _ticker_ sono per quando
// vuoi fare qualcosa ripetutamente a intervalli regolari.
// Ecco un esempio di un ticker che fa tick periodicamente
// finché non lo fermiamo.

package main

import (
	"fmt"
	"time"
)

func main() {

	// I ticker usano un meccanismo simile ai timer: un
	// canale a cui vengono inviati valori. Qui useremo il
	// `select` built-in sul canale per aspettare i
	// valori non appena arrivano ogni 500ms.
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// I ticker possono essere fermati come i timer. Una volta che un ticker
	// è fermato non riceverà più valori sul suo
	// canale. Fermeremo il nostro dopo 1600ms.
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
