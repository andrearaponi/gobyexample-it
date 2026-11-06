// A volte vorremmo che i nostri programmi Go gestiscano intelligentemente
// i [segnali Unix](https://en.wikipedia.org/wiki/Unix_signal).
// Per esempio, potremmo volere che un server si spenga
// gracefully quando riceve un `SIGTERM`, o che uno strumento
// da riga di comando smetta di processare input se riceve un `SIGINT`.
// Ecco un modo moderno di gestire i segnali usando i context.

package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
)

func main() {
	// `signal.NotifyContext` restituisce un context che viene
	// cancellato quando arriva uno dei segnali elencati.
	ctx, stop := signal.NotifyContext(
		context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Il programma aspetterà qui fino a quando non riceve
	// uno dei segnali configurati.
	fmt.Println("awaiting signal")
	<-ctx.Done()

	// `context.Cause` riporta il motivo per cui il context è
	// stato cancellato. Per una cancellazione causata da un
	// segnale, questo include il valore del segnale.
	fmt.Println()
	fmt.Println(context.Cause(ctx))
	fmt.Println("exiting")
}
