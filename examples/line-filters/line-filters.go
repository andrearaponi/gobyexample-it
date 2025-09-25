// Un _line filter_ è un tipo comune di programma che legge
// input su stdin, lo elabora, e poi stampa qualche
// risultato derivato su stdout. `grep` e `sed` sono line filter
// comuni.

// Ecco un esempio di line filter in Go che scrive una
// versione maiuscola di tutto il testo di input. Puoi usare questo
// pattern per scrivere i tuoi line filter Go.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Avvolgere l'`os.Stdin` non bufferizzato con uno scanner
	// bufferizzato ci dà un comodo metodo `Scan` che
	// avanza lo scanner al prossimo token; che è
	// la prossima linea nello scanner di default.
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// `Text` restituisce il token corrente, qui la prossima linea,
		// dall'input.
		ucl := strings.ToUpper(scanner.Text())

		// Scrivi la linea in maiuscolo.
		fmt.Println(ucl)
	}

	// Controlla per errori durante `Scan`. La fine del file è
	// attesa e non riportata da `Scan` come errore.
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
