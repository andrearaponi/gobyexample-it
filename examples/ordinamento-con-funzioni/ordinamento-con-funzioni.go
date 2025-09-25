// A volte vorremo ordinare una collezione per qualcosa
// diverso dal suo ordine naturale. Per esempio, supponiamo di
// voler ordinare le stringhe per lunghezza invece che
// alfabeticamente. Ecco un esempio di ordinamento personalizzato
// in Go.

package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fruits := []string{"peach", "banana", "kiwi"}

	// Implementiamo una funzione di confronto per le lunghezze
	// delle stringhe. `cmp.Compare` è utile per questo.
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	// Ora possiamo chiamare `slices.SortFunc` con questa funzione
	// di confronto personalizzata per ordinare `fruits` per lunghezza del nome.
	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	// Possiamo usare la stessa tecnica per ordinare uno slice di
	// valori che non sono tipi built-in.
	type Person struct {
		name string
		age  int
	}

	people := []Person{
		Person{name: "Jax", age: 37},
		Person{name: "TJ", age: 25},
		Person{name: "Alex", age: 72},
	}

	// Ordiniamo `people` per età usando `slices.SortFunc`.
	//
	// Nota: se la struct `Person` è grande,
	// potresti voler che lo slice contenga `*Person` invece
	// e adattare la funzione di ordinamento di conseguenza. In caso
	// di dubbio, [fai benchmark](testing-and-benchmarking)!
	slices.SortFunc(people,
		func(a, b Person) int {
			return cmp.Compare(a.age, b.age)
		})
	fmt.Println(people)
}
