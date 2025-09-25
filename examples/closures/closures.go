// Go supporta le [_funzioni anonime_](https://en.wikipedia.org/wiki/Anonymous_function),
// che possono formare <a href="https://en.wikipedia.org/wiki/Closure_(computer_science)"><em>closure</em></a>.
// Le funzioni anonime sono utili quando vuoi definire
// una funzione inline senza doverla nominare.

package main

import "fmt"

// Questa funzione `intSeq` restituisce un'altra funzione, che
// definiamo anonimamente nel corpo di `intSeq`. La
// funzione restituita _cattura_ la variabile `i` per
// formare una closure.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// Chiamiamo `intSeq`, assegnando il risultato (una funzione)
	// a `nextInt`. Questo valore di funzione cattura il suo
	// proprio valore `i`, che sarà aggiornato ogni volta
	// che chiamiamo `nextInt`.
	nextInt := intSeq()

	// Vedi l'effetto della closure chiamando `nextInt`
	// alcune volte.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// Per confermare che lo stato è unico per quella
	// particolare funzione, creiamo e testiamo una nuova.
	newInts := intSeq()
	fmt.Println(newInts())
}
