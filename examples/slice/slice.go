// Le _slice_ sono un tipo di dato importante in Go, fornendo
// un'interfaccia più potente per le sequenze rispetto agli array.

package main

import (
	"fmt"
	"slices"
)

func main() {

	// A differenza degli array, le slice sono tipizzate solo in base agli
	// elementi che contengono (non al numero di elementi).
	// Una slice non inizializzata è uguale a nil e ha
	// lunghezza 0.
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// Per creare una slice con lunghezza diversa da zero, usa
	// la funzione builtin `make`. Qui creiamo una slice di
	// `string` di lunghezza `3` (inizialmente con valore zero).
	// Di default la capacità di una nuova slice è uguale alla sua
	// lunghezza; se sappiamo che la slice crescerà in anticipo,
	// è possibile passare una capacità esplicitamente
	// come parametro aggiuntivo a `make`.
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// Possiamo impostare e ottenere valori proprio come con gli array.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// `len` restituisce la lunghezza della slice come previsto.
	fmt.Println("len:", len(s))

	// Oltre a queste operazioni di base, le slice
	// supportano diverse altre operazioni che le rendono più ricche degli
	// array. Una è la funzione builtin `append`, che
	// restituisce una slice contenente uno o più nuovi valori.
	// Nota che dobbiamo accettare un valore di ritorno da
	// `append` poiché potremmo ottenere un nuovo valore di slice.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Le slice possono anche essere copiate con `copy`. Qui creiamo una
	// slice vuota `c` della stessa lunghezza di `s` e copiamo
	// in `c` da `s`.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Le slice supportano un operatore "slice" con la sintassi
	// `slice[low:high]`. Per esempio, questo ottiene una slice
	// degli elementi `s[2]`, `s[3]`, e `s[4]`.
	l := s[2:5]
	fmt.Println("sl1:", l)

	// Questo crea una slice fino a (ma escludendo) `s[5]`.
	l = s[:5]
	fmt.Println("sl2:", l)

	// E questo crea una slice da (e includendo) `s[2]`.
	l = s[2:]
	fmt.Println("sl3:", l)

	// Possiamo dichiarare e inizializzare una variabile per slice
	// anche in una singola riga.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Il pacchetto `slices` contiene diverse funzioni di utilità
	// utili per le slice.
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// Le slice possono essere composte in strutture dati
	// multidimensionali. La lunghezza delle slice interne può
	// variare, a differenza degli array multidimensionali.
	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
