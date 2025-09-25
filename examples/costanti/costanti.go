// Go supporta _costanti_ di caratteri, stringhe, booleani
// e valori numerici.

package main

import (
	"fmt"
	"math"
)

// `const` dichiara un valore costante.
const s string = "constant"

func main() {
	fmt.Println(s)

	// Un'istruzione `const` può apparire ovunque possa
	// apparire un'istruzione `var`.
	const n = 500000000

	// Le espressioni costanti eseguono l'aritmetica con
	// precisione arbitraria.
	const d = 3e20 / n
	fmt.Println(d)

	// Una costante numerica non ha tipo finché non gliene
	// viene assegnato uno, ad esempio tramite una conversione esplicita.
	fmt.Println(int64(d))

	// A un numero può essere assegnato un tipo usandolo in un
	// contesto che ne richiede uno, come un'assegnazione di
	// variabile o una chiamata di funzione. Per esempio, qui
	// `math.Sin` si aspetta un `float64`.
	fmt.Println(math.Sin(n))
}
