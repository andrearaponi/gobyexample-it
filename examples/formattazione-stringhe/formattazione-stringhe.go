// Go offre un eccellente supporto per la formattazione delle stringhe nella
// tradizione `printf`. Ecco alcuni esempi di
// attività comuni di formattazione delle stringhe.

package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	// Go offre diversi "verbi" di stampa progettati per
	// formattare valori Go generali. Per esempio, questo stampa
	// un'istanza della nostra struct `point`.
	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p)

	// Se il valore è una struct, la variante `%+v`
	// includerà i nomi dei campi della struct.
	fmt.Printf("struct2: %+v\n", p)

	// La variante `%#v` stampa una rappresentazione della sintassi Go
	// del valore, cioè il frammento di codice sorgente che
	// produrrebbe quel valore.
	fmt.Printf("struct3: %#v\n", p)

	// Per stampare il tipo di un valore, usa `%T`.
	fmt.Printf("type: %T\n", p)

	// La formattazione dei booleani è diretta.
	fmt.Printf("bool: %t\n", true)

	// Ci sono molte opzioni per formattare gli interi.
	// Usa `%d` per la formattazione standard, base-10.
	fmt.Printf("int: %d\n", 123)

	// Questo stampa una rappresentazione binaria.
	fmt.Printf("bin: %b\n", 14)

	// Questo stampa il carattere corrispondente
	// all'intero dato.
	fmt.Printf("char: %c\n", 33)

	// `%x` fornisce la codifica esadecimale.
	fmt.Printf("hex: %x\n", 456)

	// Ci sono anche diverse opzioni di formattazione per
	// i float. Per la formattazione decimale di base usa `%f`.
	fmt.Printf("float1: %f\n", 78.9)

	// `%e` e `%E` formattano il float in (versioni leggermente
	// diverse della) notazione scientifica.
	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float3: %E\n", 123400000.0)

	// Per la stampa di base delle stringhe usa `%s`.
	fmt.Printf("str1: %s\n", "\"string\"")

	// Per racchiudere le stringhe tra virgolette doppie come nel sorgente Go, usa `%q`.
	fmt.Printf("str2: %q\n", "\"string\"")

	// Come con gli interi visti prima, `%x` rende
	// la stringa in base-16, con due caratteri di output
	// per byte di input.
	fmt.Printf("str3: %x\n", "hex this")

	// Per stampare una rappresentazione di un puntatore, usa `%p`.
	fmt.Printf("pointer: %p\n", &p)

	// Quando formatti i numeri spesso vorrai
	// controllare la larghezza e la precisione della figura
	// risultante. Per specificare la larghezza di un intero, usa un
	// numero dopo il `%` nel verbo. Per default il
	// risultato sarà allineato a destra e riempito con
	// spazi.
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	// Puoi anche specificare la larghezza dei float stampati,
	// anche se di solito vorrai anche limitare la
	// precisione decimale allo stesso tempo con la
	// sintassi width.precision.
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	// Per allineare a sinistra, usa il flag `-`.
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// Potresti anche voler controllare la larghezza quando formatti
	// le stringhe, specialmente per assicurarti che si allineino in
	// output simili a tabelle. Per larghezza base allineata a destra.
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

	// Per allineare a sinistra usa il flag `-` come con i numeri.
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	// Finora abbiamo visto `Printf`, che stampa la
	// stringa formattata su `os.Stdout`. `Sprintf` formatta
	// e restituisce una stringa senza stamparla da nessuna parte.
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	// Puoi formattare+stampare su `io.Writers` diversi da
	// `os.Stdout` usando `Fprintf`.
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
