// `for` è l'unico costrutto di ciclo di Go. Ecco
// alcuni tipi di base di cicli `for`.

package main

import "fmt"

func main() {

	// Il tipo più semplice, con una singola condizione.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Un classico ciclo `for` con inizializzazione/condizione/incremento.
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// Un altro modo per realizzare l'iterazione di base
	// "fai questo N volte" è usare `range` su un intero.
	for i := range 3 {
		fmt.Println("range", i)
	}

	// `for` senza condizione esegue il ciclo ripetutamente
	// finché non esci con `break` dal ciclo o con `return`
	// dalla funzione che lo contiene.
	for {
		fmt.Println("loop")
		break
	}

	// Puoi anche usare `continue` per passare alla prossima
	// iterazione del ciclo.
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
