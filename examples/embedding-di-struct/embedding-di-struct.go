// Go supporta l'_embedding_ di struct e interfacce
// per esprimere una _composizione_ di tipi più fluida.
// Questo non deve essere confuso con [`//go:embed`](embed-directive) che è
// una direttiva go introdotta in Go versione 1.16+ per incorporare
// file e cartelle nel binario dell'applicazione.

package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// Un `container` _incorpora_ un `base`. Un embedding assomiglia
// a un campo senza nome.
type container struct {
	base
	str string
}

func main() {

	// Quando creiamo struct con letterali, dobbiamo
	// inizializzare l'embedding esplicitamente; qui il
	// tipo incorporato funge da nome del campo.
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// Possiamo accedere ai campi di base direttamente su `co`,
	// ad esempio `co.num`.
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// In alternativa, possiamo specificare il percorso completo usando
	// il nome del tipo incorporato.
	fmt.Println("also num:", co.base.num)

	// Poiché `container` incorpora `base`, i metodi di
	// `base` diventano anche metodi di un `container`. Qui
	// invochiamo un metodo che è stato incorporato da `base`
	// direttamente su `co`.
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// L'embedding di struct con metodi può essere usato per conferire
	// implementazioni di interfacce ad altre struct. Qui
	// vediamo che un `container` ora implementa l'
	// interfaccia `describer` perché incorpora `base`.
	var d describer = co
	fmt.Println("describer:", d.describe())
}
