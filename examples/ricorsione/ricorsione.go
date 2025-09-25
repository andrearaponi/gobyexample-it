// Go supporta le
// <a href="https://en.wikipedia.org/wiki/Recursion_(computer_science)"><em>funzioni ricorsive</em></a>.
// Ecco un esempio classico.

package main

import "fmt"

// Questa funzione `fact` chiama se stessa fino a raggiungere il
// caso base di `fact(0)`.
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	// Anche le funzioni anonime possono essere ricorsive, ma questo richiede
	// di dichiarare esplicitamente una variabile con `var` per memorizzare
	// la funzione prima che sia definita.
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		// Poiché `fib` è stata precedentemente dichiarata in `main`, Go
		// sa quale funzione chiamare con `fib` qui.
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
