// Go ha vari tipi di valori tra cui stringhe,
// interi, float, booleani, ecc. Ecco alcuni
// esempi di base.

package main

import "fmt"

func main() {

	// Stringhe, che possono essere concatenate con `+`.
	fmt.Println("go" + "lang")

	// Interi e float.
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Booleani, con gli operatori booleani come ti aspetteresti.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
