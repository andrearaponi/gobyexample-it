// Usa `os.Exit` per uscire immediatamente con un
// status dato.

package main

import (
	"fmt"
	"os"
)

func main() {

	// I `defer` _non_ verranno eseguiti quando si usa `os.Exit`,
	// quindi questo `fmt.Println` non verrà mai chiamato.
	defer fmt.Println("!")

	// Esci con status 3.
	os.Exit(3)
}

// Nota che a differenza di es. C, Go non usa un valore
// di ritorno intero da `main` per indicare lo status di uscita.
// Se vuoi uscire con uno status diverso da zero dovresti
// usare `os.Exit`.
