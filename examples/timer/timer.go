// Spesso vogliamo eseguire codice Go a un certo punto nel
// futuro, o ripetutamente a un certo intervallo. Le funzionalità
// built-in _timer_ e _ticker_ di Go rendono entrambi questi
// compiti facili. Vedremo prima i timer e poi
// i [ticker](tickers).

package main

import (
	"fmt"
	"time"
)

func main() {

	// I timer rappresentano un singolo evento nel futuro. Dici
	// al timer quanto vuoi aspettare, ed esso fornisce un
	// canale che sarà notificato in quel momento. Questo timer
	// aspetterà 2 secondi.
	timer1 := time.NewTimer(2 * time.Second)

	// Il `<-timer1.C` si blocca sul canale `C` del timer
	// finché non invia un valore indicando che il timer
	// è scattato.
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// Se volevi solo aspettare, avresti potuto usare
	// `time.Sleep`. Un motivo per cui un timer può essere utile è
	// che puoi cancellare il timer prima che scatti.
	// Ecco un esempio di questo.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// Diamo al `timer2` abbastanza tempo per scattare, se mai
	// dovesse farlo, per mostrare che è effettivamente fermato.
	time.Sleep(2 * time.Second)
}
