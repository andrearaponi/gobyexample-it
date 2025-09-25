// In Go, le _variabili_ sono dichiarate esplicitamente e usate dal
// compilatore per es. controllare la correttezza dei tipi nelle
// chiamate di funzione.

package main

import "fmt"

func main() {

	// `var` dichiara 1 o più variabili.
	var a = "initial"
	fmt.Println(a)

	// Puoi dichiarare più variabili contemporaneamente.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go dedurrà il tipo delle variabili inizializzate.
	var d = true
	fmt.Println(d)

	// Le variabili dichiarate senza una corrispondente
	// inizializzazione hanno un _valore zero_. Per esempio, il
	// valore zero per un `int` è `0`.
	var e int
	fmt.Println(e)

	// La sintassi `:=` è una scorciatoia per dichiarare e
	// inizializzare una variabile, es. per
	// `var f string = "apple"` in questo caso.
	// Questa sintassi è disponibile solo all'interno delle funzioni.
	f := "apple"
	fmt.Println(f)
}
