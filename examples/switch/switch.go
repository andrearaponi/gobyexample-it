// Le _istruzioni switch_ esprimono condizionali su molti
// rami.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Ecco uno `switch` di base.
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// Puoi usare le virgole per separare più espressioni
	// nella stessa istruzione `case`. In questo esempio usiamo
	// anche il caso `default` opzionale.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// Uno `switch` senza espressione è un modo alternativo
	// per esprimere la logica if/else. Qui mostriamo anche come
	// le espressioni `case` possono essere non costanti.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// Uno `switch` di tipo confronta i tipi invece dei valori. Puoi
	// usarlo per scoprire il tipo di un valore di
	// interfaccia. In questo esempio, la variabile `t` avrà il
	// tipo corrispondente alla sua clausola.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
