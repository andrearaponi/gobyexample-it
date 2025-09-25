// [_I flag da linea di comando_](https://en.wikipedia.org/wiki/Command-line_interface#Command-line_option)
// sono un modo comune per specificare opzioni
// per programmi da linea di comando. Per esempio,
// in `wc -l` il `-l` è un flag da linea di comando.

package main

// Go fornisce un pacchetto `flag` che supporta
// il parsing di base dei flag da linea di comando.
// Useremo questo pacchetto per implementare
// il nostro programma di esempio.
import (
	"flag"
	"fmt"
)

func main() {

	// Dichiarazioni di flag di base sono disponibili
	// per opzioni string, integer e boolean. Qui
	// dichiariamo un flag string `word` con valore
	// di default `"foo"` e una breve descrizione.
	// Questa funzione `flag.String` restituisce un
	// puntatore string; vedremo come usarlo sotto.
	wordPtr := flag.String("word", "foo", "a string")

	// Questo dichiara i flag `numb` e `fork`, usando
	// un approccio simile al flag `word`.
	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	// È anche possibile dichiarare un'opzione che usa
	// una var esistente dichiarata altrove nel programma.
	// Nota che dobbiamo passare un puntatore alla
	// funzione di dichiarazione del flag.
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// Una volta dichiarati tutti i flag, chiama `flag.Parse()`
	// per eseguire il parsing da linea di comando.
	flag.Parse()

	// Qui stamperemo le opzioni parsate e qualsiasi
	// argomento posizionale finale. Nota che dobbiamo
	// dereferenziare i puntatori con es. `*wordPtr`
	// per ottenere i valori effettivi delle opzioni.
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
