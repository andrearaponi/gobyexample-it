// Una stringa Go è una slice di sola lettura di byte. Il linguaggio
// e la libreria standard trattano le stringhe in modo speciale - come
// contenitori di testo codificato in [UTF-8](https://en.wikipedia.org/wiki/UTF-8).
// In altri linguaggi, le stringhe sono costituite da "caratteri".
// In Go, il concetto di carattere è chiamato `rune` - è
// un intero che rappresenta un punto di codice Unicode.
// [Questo post del blog Go](https://go.dev/blog/strings) è una buona
// introduzione all'argomento.

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// `s` è una `string` a cui è assegnato un valore letterale
	// che rappresenta la parola "ciao" in lingua
	// thailandese. I letterali stringa Go sono testo
	// codificato UTF-8.
	const s = "สวัสดี"

	// Poiché le stringhe sono equivalenti a `[]byte`, questo
	// produrrà la lunghezza dei byte grezzi memorizzati all'interno.
	fmt.Println("Len:", len(s))

	// L'indicizzazione di una stringa produce i valori di byte grezzi a
	// ogni indice. Questo ciclo genera i valori esadecimali di tutti
	// i byte che costituiscono i punti di codice in `s`.
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// Per contare quante _rune_ ci sono in una stringa, possiamo usare
	// il pacchetto `utf8`. Nota che il tempo di esecuzione di
	// `RuneCountInString` dipende dalla dimensione della stringa,
	// perché deve decodificare ogni rune UTF-8 sequenzialmente.
	// Alcuni caratteri thailandesi sono rappresentati da punti di codice UTF-8
	// che possono estendersi su più byte, quindi il risultato di questo conteggio
	// potrebbe essere sorprendente.
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// Un ciclo `range` gestisce le stringhe in modo speciale e decodifica
	// ogni `rune` insieme al suo offset nella stringa.
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// Possiamo ottenere la stessa iterazione usando
	// esplicitamente la funzione `utf8.DecodeRuneInString`.
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		// Questo dimostra il passaggio di un valore `rune` a una funzione.
		examineRune(runeValue)
	}
}

func examineRune(r rune) {

	// I valori racchiusi tra virgolette singole sono _letterali rune_. Possiamo
	// confrontare un valore `rune` con un letterale rune direttamente.
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
