// Le _mappe_ sono il [tipo di dato associativo](https://en.wikipedia.org/wiki/Associative_array) built-in di Go
// (a volte chiamate _hash_ o _dizionari_ in altri linguaggi).

package main

import (
	"fmt"
	"maps"
)

func main() {

	// Per creare una mappa vuota, usa la funzione builtin `make`:
	// `make(map[tipo-chiave]tipo-valore)`.
	m := make(map[string]int)

	// Imposta coppie chiave/valore usando la tipica sintassi
	// `nome[chiave] = valore`.
	m["k1"] = 7
	m["k2"] = 13

	// Stampare una mappa con ad es. `fmt.Println` mostrerà tutte
	// le sue coppie chiave/valore.
	fmt.Println("map:", m)

	// Ottieni un valore per una chiave con `nome[chiave]`.
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// Se la chiave non esiste, viene restituito il
	// [valore zero](https://go.dev/ref/spec#The_zero_value) del
	// tipo del valore.
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	// La funzione builtin `len` restituisce il numero di coppie
	// chiave/valore quando chiamata su una mappa.
	fmt.Println("len:", len(m))

	// La funzione builtin `delete` rimuove coppie chiave/valore da
	// una mappa.
	delete(m, "k2")
	fmt.Println("map:", m)

	// Per rimuovere *tutte* le coppie chiave/valore da una mappa, usa
	// la funzione builtin `clear`.
	clear(m)
	fmt.Println("map:", m)

	// Il secondo valore di ritorno opzionale quando si ottiene un
	// valore da una mappa indica se la chiave era presente
	// nella mappa. Questo può essere usato per distinguere
	// tra chiavi mancanti e chiavi con valori zero
	// come `0` o `""`. Qui non avevamo bisogno del valore
	// stesso, quindi l'abbiamo ignorato con l'_identificatore vuoto_
	// `_`.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// Puoi anche dichiarare e inizializzare una nuova mappa
	// nella stessa riga con questa sintassi.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// Il pacchetto `maps` contiene diverse funzioni di utilità
	// utili per le mappe.
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
