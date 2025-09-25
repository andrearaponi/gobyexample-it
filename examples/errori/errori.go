// In Go è idiomatico comunicare gli errori tramite un
// valore di ritorno esplicito e separato. Questo si contrappone
// alle eccezioni usate in linguaggi come Java e Ruby e al
// singolo valore di ritorno che, in C, talvolta funge sia da
// risultato che da errore. L'approccio di Go rende immediato
// capire quali funzioni restituiscono errori e gestirli
// utilizzando gli stessi costrutti del linguaggio impiegati
// anche per altre operazioni non legate agli errori.
//
// Vedi la documentazione del [pacchetto errors](https://pkg.go.dev/errors)
// e [questo blog post](https://go.dev/blog/go1.13-errors) per dettagli
// aggiuntivi.

package main

import (
	"errors"
	"fmt"
)

// Per convenzione, gli errori sono l'ultimo valore di ritorno e
// hanno tipo `error`, un'interfaccia built-in.
func f(arg int) (int, error) {
	if arg == 42 {
		// `errors.New` costruisce un valore `error` di base
		// con il messaggio di errore dato.
		return -1, errors.New("can't work with 42")
	}

	// Un valore `nil` nella posizione dell'errore indica che
	// non c'è stato alcun errore.
	return arg + 3, nil
}

// Un errore sentinella è una variabile predichiarata che è usata per
// indicare una condizione di errore specifica.
var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {

		// Possiamo avvolgere gli errori con errori di livello superiore per aggiungere
		// contesto. Il modo più semplice per farlo è con il
		// verbo `%w` in `fmt.Errorf`. Gli errori avvolti
		// creano una catena logica (A avvolge B, che avvolge C, ecc.)
		// che può essere interrogata con funzioni come `errors.Is`
		// e `errors.As`.
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func main() {
	for _, i := range []int{7, 42} {

		// È comune usare un controllo di errore inline nella riga
		// `if`.
		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {

			// `errors.Is` controlla che un dato errore (o qualsiasi errore nella sua catena)
			// corrisponda a un valore di errore specifico. Questo è particolarmente utile con errori avvolti o
			// annidati, permettendoti di identificare tipi di errore specifici o errori
			// sentinella in una catena di errori.
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark.")
			} else {
				fmt.Printf("unknown error: %s\n", err)
			}
			continue
		}

		fmt.Println("Tea is ready!")
	}
}
