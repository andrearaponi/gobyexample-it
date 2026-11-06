// È possibile definire tipi di errore personalizzati
// implementando il metodo `Error()` su di essi. Ecco una
// variante dell'esempio precedente che usa un tipo personalizzato
// per rappresentare esplicitamente un errore di argomento.

package main

import (
	"errors"
	"fmt"
)

// Un tipo di errore personalizzato solitamente ha il suffisso "Error".
type argError struct {
	arg     int
	message string
}

// Aggiungendo questo metodo `Error` facciamo sì che `argError`
// implementi l'interfaccia `error`.
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {

		// Restituiamo il nostro errore personalizzato.
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	// `errors.AsType` è una versione più avanzata di `errors.Is`.
	// Controlla che un dato errore (o qualsiasi errore nella sua catena)
	// corrisponda a un tipo di errore specifico e lo converte in un valore
	// di quel tipo, restituendo anche `true`. Se non c'è corrispondenza,
	// il secondo valore di ritorno è `false`.
	_, err := f(42)
	if ae, ok := errors.AsType[*argError](err); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}
