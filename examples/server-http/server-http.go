// Scrivere un server HTTP di base è facile usando
// il pacchetto `net/http`.
package main

import (
	"fmt"
	"net/http"
)

// Un concetto fondamentale nei server `net/http` sono
// gli *handler*. Un handler è un oggetto che implementa
// l'interfaccia `http.Handler`. Un modo comune per scrivere
// un handler è usare l'adapter `http.HandlerFunc`
// su funzioni con la signature appropriata.
func hello(w http.ResponseWriter, req *http.Request) {

	// Le funzioni che fungono da handler prendono un
	// `http.ResponseWriter` e un `http.Request` come
	// argomenti. Il response writer è usato per riempire
	// la risposta HTTP. Qui la nostra semplice risposta
	// è solo "hello\n".
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	// Questo handler fa qualcosa di un po' più
	// sofisticato leggendo tutti gli header della richiesta
	// HTTP e stampandoli nel body della risposta.
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	// Registriamo i nostri handler sulle route del server usando
	// la funzione di convenienza `http.HandleFunc`. Configura
	// il *router di default* nel pacchetto `net/http` e
	// prende una funzione come argomento.
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// Infine, chiamiamo `ListenAndServe` con la porta
	// e un handler. `nil` dice di usare il router
	// di default che abbiamo appena configurato.
	http.ListenAndServe(":8090", nil)
}
