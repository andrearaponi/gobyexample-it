// Go ha il supporto built-in per _valori di ritorno multipli_.
// Questa caratteristica è usata spesso nel Go idiomatico, per esempio
// per restituire sia il risultato che i valori di errore da una funzione.

package main

import "fmt"

// Il `(int, int)` in questa firma di funzione mostra che
// la funzione restituisce 2 `int`.
func vals() (int, int) {
	return 3, 7
}

func main() {

	// Qui usiamo i 2 diversi valori di ritorno dalla
	// chiamata con _assegnazione multipla_.
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// Se vuoi solo un sottoinsieme dei valori restituiti,
	// usa l'identificatore vuoto `_`.
	_, c := vals()
	fmt.Println(c)
}
