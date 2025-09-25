// A partire dalla versione 1.23, Go ha aggiunto il supporto per
// gli [iteratori](https://go.dev/blog/range-functions),
// che ci permettono di fare range su praticamente qualsiasi cosa!

package main

import (
	"fmt"
	"iter"
	"slices"
)

// Diamo un'occhiata di nuovo al tipo `List` dall'
// [esempio precedente](generici). In quell'esempio
// avevamo un metodo `AllElements` che restituiva una slice
// di tutti gli elementi nella lista. Con gli iteratori Go, possiamo
// farlo meglio - come mostrato sotto.
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// All restituisce un _iteratore_, che in Go è una funzione
// con una [signature speciale](https://pkg.go.dev/iter#Seq).
func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		// La funzione iteratore prende un'altra funzione come
		// parametro, chiamata `yield` per convenzione (ma
		// il nome può essere arbitrario). Chiamerà `yield` per
		// ogni elemento su cui vogliamo iterare, e noterà il valore
		// di ritorno di `yield` per una potenziale terminazione anticipata.
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

// L'iterazione non richiede una struttura dati sottostante,
// e non deve nemmeno essere finita! Ecco una funzione
// che restituisce un iteratore sui numeri di Fibonacci: continua
// a funzionare finché `yield` continua a restituire `true`.
func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	// Poiché `List.All` restituisce un iteratore, possiamo usarlo
	// in un ciclo `range` normale.
	for e := range lst.All() {
		fmt.Println(e)
	}

	// Pacchetti come [slices](https://pkg.go.dev/slices) hanno
	// un numero di funzioni utili per lavorare con gli iteratori.
	// Per esempio, `Collect` prende qualsiasi iteratore e raccoglie
	// tutti i suoi valori in una slice.
	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	for n := range genFib() {

		// Una volta che il ciclo incontra `break` o un return anticipato, la funzione `yield`
		// passata all'iteratore restituirà `false`.
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}
