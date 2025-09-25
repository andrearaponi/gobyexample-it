// Nell'esempio precedente abbiamo visto come configurare un semplice
// [server HTTP](server-http). I server HTTP sono utili per
// dimostrare l'uso di `context.Context` per
// controllare la cancellazione. Un `Context` trasporta scadenze,
// segnali di cancellazione e altri valori legati alla richiesta
// attraverso i confini delle API e le goroutine.
package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	// Un `context.Context` è creato per ogni richiesta dal
	// meccanismo `net/http`, ed è disponibile con
	// il metodo `Context()`.
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	// Aspetta alcuni secondi prima di inviare una risposta al
	// client. Questo potrebbe simulare del lavoro che il server sta
	// facendo. Mentre lavora, tieni d'occhio il canale
	// `Done()` del context per un segnale che dovremmo cancellare
	// il lavoro e tornare il prima possibile.
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		// Il metodo `Err()` del context restituisce un errore
		// che spiega perché il canale `Done()` è stato
		// chiuso.
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {

	// Come prima, registriamo il nostro handler sulla route
	// "/hello" e iniziamo a servire.
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
