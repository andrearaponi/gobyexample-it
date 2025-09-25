// Il pacchetto `slices` di Go implementa l'ordinamento per i tipi built-in
// e i tipi definiti dall'utente. Prima vedremo l'ordinamento per
// i tipi built-in.

package main

import (
	"fmt"
	"slices"
)

func main() {

	// Le funzioni di ordinamento sono generiche e funzionano per qualsiasi
	// tipo built-in _ordinato_. Per un elenco dei tipi ordinati,
	// vedi [cmp.Ordered](https://pkg.go.dev/cmp#Ordered).
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)

	// Un esempio di ordinamento di `int`.
	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints:   ", ints)

	// Possiamo anche usare il pacchetto `slices` per verificare se
	// uno slice è già ordinato.
	s := slices.IsSorted(ints)
	fmt.Println("Sorted: ", s)
}
