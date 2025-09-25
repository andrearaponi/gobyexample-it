// [_Gli argomenti da linea di comando_](https://en.wikipedia.org/wiki/Command-line_interface#Arguments)
// sono un modo comune per parametrizzare
// l'esecuzione dei programmi. Per esempio,
// `go run hello.go` usa argomenti.

package main

import (
	"fmt"
	"os"
)

func main() {

	// `os.Args` fornisce accesso agli argomenti
	// grezzi. Il primo valore è il percorso
	// del programma, `os.Args[1:]` contiene
	// gli argomenti del programma.
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// Argomenti individuali con indicizzazione.
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
