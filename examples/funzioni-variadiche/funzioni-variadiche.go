// Le [_funzioni variadiche_](https://en.wikipedia.org/wiki/Variadic_function)
// possono essere chiamate con qualsiasi numero di argomenti finali.
// Per esempio, `fmt.Println` è una comune funzione
// variadica.

package main

import "fmt"

// Ecco una funzione che prenderà un numero arbitrario
// di `int` come argomenti.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	// All'interno della funzione, il tipo di `nums` è
	// equivalente a `[]int`. Possiamo chiamare `len(nums)`,
	// iterare su di esso con `range`, ecc.
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	// Le funzioni variadiche possono essere chiamate nel modo usuale
	// con argomenti individuali.
	sum(1, 2)
	sum(1, 2, 3)

	// Se hai già più argomenti in una slice,
	// applicali a una funzione variadica usando
	// `func(slice...)` così.
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
