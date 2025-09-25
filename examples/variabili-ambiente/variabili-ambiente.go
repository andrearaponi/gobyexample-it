// Le [variabili d'ambiente](https://en.wikipedia.org/wiki/Environment_variable)
// sono un meccanismo universale per [trasmettere informazioni
// di configurazione ai programmi Unix](https://www.12factor.net/config).
// Vediamo come impostare, ottenere ed elencare le variabili d'ambiente.

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// Per impostare una coppia chiave/valore, usa `os.Setenv`.
	// Per ottenere un valore per una chiave, usa `os.Getenv`.
	// Questo restituirà una stringa vuota se la chiave
	// non è presente nell'ambiente.
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	// Usa `os.Environ` per elencare tutte le coppie
	// chiave/valore nell'ambiente. Questo restituisce
	// uno slice di stringhe nella forma `KEY=value`.
	// Puoi usare `strings.SplitN` per ottenere la chiave
	// e il valore. Qui stampiamo tutte le chiavi.
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
