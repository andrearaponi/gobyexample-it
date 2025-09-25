// Go rende possibile _recuperare_ da un panic, usando
// la funzione built-in `recover`. Un `recover` può
// fermare un `panic` dall'interrompere il programma e lasciarlo
// continuare con l'esecuzione invece.

// Un esempio di dove questo può essere utile: un server
// non vorrebbe crashare se una delle connessioni client
// presenta un errore critico. Invece, il server vorrebbe
// chiudere quella connessione e continuare a servire
// altri client. In effetti, questo è quello che fa Go's `net/http`
// di default per i server HTTP.

package main

import "fmt"

// Questa funzione fa panic.
func mayPanic() {
	panic("a problem")
}

func main() {
	// `recover` deve essere chiamato all'interno di una funzione deferred.
	// Quando la funzione che la racchiude fa panic, il defer si
	// attiverà e una chiamata a `recover` al suo interno catturerà
	// il panic.
	defer func() {
		if r := recover(); r != nil {
			// Il valore di ritorno di `recover` è l'errore sollevato nella
			// chiamata a `panic`.
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	// Questo codice non verrà eseguito, perché `mayPanic` fa panic.
	// L'esecuzione di `main` si ferma al punto del
	// panic e riprende nella closure deferred.
	fmt.Println("After mayPanic()")
}
