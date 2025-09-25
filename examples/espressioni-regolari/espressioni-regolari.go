// Go offre supporto integrato per le [espressioni regolari](https://en.wikipedia.org/wiki/Regular_expression).
// Ecco alcuni esempi di operazioni comuni relative alle regexp
// in Go.

package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	// Questo testa se un pattern corrisponde a una stringa.
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// Sopra abbiamo usato un pattern string direttamente, ma per
	// altre operazioni regexp dovrai compilare (`Compile`) un
	// struct `Regexp` ottimizzato.
	r, _ := regexp.Compile("p([a-z]+)ch")

	// Molti metodi sono disponibili su questi struct. Ecco
	// un test di corrispondenza come abbiamo visto prima.
	fmt.Println(r.MatchString("peach"))

	// Questo trova la corrispondenza per la regexp.
	fmt.Println(r.FindString("peach punch"))

	// Questo trova anche la prima corrispondenza ma restituisce
	// gli indici di inizio e fine per la corrispondenza invece del
	// testo corrispondente.
	fmt.Println("idx:", r.FindStringIndex("peach punch"))

	// Le varianti `Submatch` includono informazioni sia sulle
	// corrispondenze dell'intero pattern che sui submatch
	// all'interno di tali corrispondenze. Per esempio questo restituirà
	// informazioni sia per `p([a-z]+)ch` che per `([a-z]+)`.
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// Similmente questo restituirà informazioni sugli
	// indici di corrispondenze e submatch.
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// Le varianti `All` di queste funzioni si applicano a tutte
	// le corrispondenze nell'input, non solo alla prima. Per
	// esempio per trovare tutte le corrispondenze per una regexp.
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// Queste varianti `All` sono disponibili anche per le altre
	// funzioni che abbiamo visto sopra.
	fmt.Println("all:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	// Fornire un intero non-negativo come secondo
	// argomento a queste funzioni limiterà il numero
	// di corrispondenze.
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// I nostri esempi sopra avevano argomenti string e usavano
	// nomi come `MatchString`. Possiamo anche fornire
	// argomenti `[]byte` e omettere `String` dal
	// nome della funzione.
	fmt.Println(r.Match([]byte("peach")))

	// Quando crei variabili globali con espressioni
	// regolari puoi usare la variante `MustCompile`
	// di `Compile`. `MustCompile` fa panic invece di
	// restituire un errore, il che lo rende più sicuro da usare per
	// variabili globali.
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)

	// Il pacchetto `regexp` può anche essere usato per sostituire
	// sottoinsiemi di stringhe con altri valori.
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// La variante `Func` ti permette di trasformare il testo
	// corrispondente con una funzione data.
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
