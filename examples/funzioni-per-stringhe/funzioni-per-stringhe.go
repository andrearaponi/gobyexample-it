// Il pacchetto `strings` della libreria standard fornisce molte
// funzioni utili relative alle stringhe. Ecco alcuni esempi
// per darti un'idea del pacchetto.

package main

import (
	"fmt"
	s "strings"
)

// Creiamo un alias di `fmt.Println` con un nome più corto dato che lo useremo
// molto qui sotto.
var p = fmt.Println

func main() {

	// Ecco un campione delle funzioni disponibili in
	// `strings`. Dato che queste sono funzioni del
	// pacchetto, non metodi sull'oggetto stringa stesso,
	// dobbiamo passare la stringa in questione come primo
	// argomento alla funzione. Puoi trovare altre
	// funzioni nella documentazione del pacchetto [`strings`](https://pkg.go.dev/strings).
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
}
