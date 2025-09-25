// A partire dalla versione 1.18, Go ha aggiunto il supporto per
// i _generici_, anche conosciuti come _parametri di tipo_.

package main

import "fmt"

// Come esempio di funzione generica, `SlicesIndex` prende
// una slice di qualsiasi tipo `comparable` e un elemento di quel
// tipo e restituisce l'indice della prima occorrenza di
// v in s, o -1 se non presente. Il vincolo `comparable`
// significa che possiamo confrontare valori di questo tipo con gli
// operatori `==` e `!=`. Per una spiegazione più approfondita
// di questa signature di tipo, vedi [questo blog post](https://go.dev/blog/deconstructing-type-parameters).
// Nota che questa funzione esiste nella libreria standard
// come [slices.Index](https://pkg.go.dev/slices#Index).
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

// Come esempio di tipo generico, `List` è una
// lista collegata singolarmente con valori di qualsiasi tipo.
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// Possiamo definire metodi sui tipi generici proprio come
// facciamo sui tipi regolari, ma dobbiamo mantenere i parametri
// di tipo al loro posto. Il tipo è `List[T]`, non `List`.
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// AllElements restituisce tutti gli elementi della List come slice.
// Nel prossimo esempio vedremo un modo più idiomatico
// di iterare su tutti gli elementi di tipi personalizzati.
func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var s = []string{"foo", "bar", "zoo"}

	// Quando invochiamo funzioni generiche, possiamo spesso fare affidamento
	// sull'_inferenza di tipo_. Nota che non dobbiamo
	// specificare i tipi per `S` e `E` quando
	// chiamiamo `SlicesIndex` - il compilatore li inferisce
	// automaticamente.
	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

	// ... anche se potremmo anche specificarli esplicitamente.
	_ = SlicesIndex[[]string, string](s, "zoo")

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())
}
