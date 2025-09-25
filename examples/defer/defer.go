// _Defer_ viene usato per assicurare che una chiamata di funzione venga
// eseguita più tardi nell'esecuzione di un programma, di solito per
// scopi di pulizia. `defer` è spesso usato dove ad esempio
// `ensure` e `finally` sarebbero usati in altri linguaggi.

package main

import (
	"fmt"
	"os"
)

// Supponiamo di voler creare un file, scriverci,
// e poi chiuderlo quando abbiamo finito. Ecco come potremmo
// farlo con `defer`.
func main() {

	// Immediatamente dopo aver ottenuto un oggetto file con
	// `createFile`, rimandiamo la chiusura di quel file
	// con `closeFile`. Questo sarà eseguito alla fine
	// della funzione che lo racchiude (`main`), dopo
	// che `writeFile` ha finito.
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	// È importante controllare gli errori quando si chiude un
	// file, anche in una funzione deferred.
	if err != nil {
		panic(err)
	}
}
