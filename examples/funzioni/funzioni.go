// Le _funzioni_ sono centrali in Go. Impareremo le
// funzioni con alcuni esempi diversi.

package main

import "fmt"

// Ecco una funzione che prende due `int` e restituisce
// la loro somma come `int`.
func plus(a int, b int) int {

	// Go richiede return espliciti, cioè non
	// restituirà automaticamente il valore dell'ultima
	// espressione.
	return a + b
}

// Quando hai più parametri consecutivi dello
// stesso tipo, puoi omettere il nome del tipo per i
// parametri dello stesso tipo fino al parametro finale che
// dichiara il tipo.
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	// Chiama una funzione proprio come ti aspetteresti, con
	// `nome(argomenti)`.
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
