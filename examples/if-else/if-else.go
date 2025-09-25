// La ramificazione con `if` ed `else` in Go è
// semplice e diretta.

package main

import "fmt"

func main() {

	// Ecco un esempio di base.
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// Puoi avere un'istruzione `if` senza else.
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// Gli operatori logici come `&&` e `||` sono spesso
	// utili nelle condizioni.
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	// Un'istruzione può precedere i condizionali; qualsiasi variabile
	// dichiarata in questa istruzione è disponibile nel ramo corrente
	// e in tutti i rami successivi.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

// Nota che non sono necessarie le parentesi attorno alle condizioni
// in Go, ma le parentesi graffe sono obbligatorie.
