// Go supporta i <em><a href="https://en.wikipedia.org/wiki/Pointer_(computer_programming)">puntatori</a></em>,
// permettendoti di passare riferimenti a valori e record
// all'interno del tuo programma.

package main

import "fmt"

// Mostreremo come funzionano i puntatori in contrasto ai valori con
// 2 funzioni: `zeroval` e `zeroptr`. `zeroval` ha un
// parametro `int`, quindi gli argomenti verranno passati ad essa per
// valore. `zeroval` riceverà una copia di `ival` distinta
// da quella nella funzione chiamante.
func zeroval(ival int) {
	ival = 0
}

// `zeroptr` al contrario ha un parametro `*int`, il che significa
// che accetta un puntatore `int`. Il codice `*iptr` nel
// corpo della funzione poi _dereferenzia_ il puntatore dal suo
// indirizzo di memoria al valore corrente a quell'indirizzo.
// Assegnare un valore a un puntatore dereferenziato cambia il
// valore all'indirizzo referenziato.
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// La sintassi `&i` fornisce l'indirizzo di memoria di `i`,
	// cioè un puntatore a `i`.
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// Anche i puntatori possono essere stampati.
	fmt.Println("pointer:", &i)
}
