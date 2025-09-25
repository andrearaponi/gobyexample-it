// _range_ itera sugli elementi in una varietà di
// strutture dati built-in. Vediamo come
// usare `range` con alcune delle strutture dati
// che abbiamo già imparato.

package main

import "fmt"

func main() {

	// Qui usiamo `range` per sommare i numeri in una slice.
	// Anche gli array funzionano così.
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// `range` su array e slice fornisce sia l'
	// indice che il valore per ogni elemento. Sopra non
	// avevamo bisogno dell'indice, quindi l'abbiamo ignorato con
	// l'identificatore blank `_`. A volte però vogliamo
	// effettivamente gli indici.
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// `range` su map itera sulle coppie chiave/valore.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// `range` può anche iterare solo sulle chiavi di una map.
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// `range` su stringhe itera sui punti di codice Unicode.
	// Il primo valore è l'indice di byte iniziale
	// della `rune` e il secondo la `rune` stessa.
	// Vedi [Strings and Runes](strings-and-runes) per maggiori
	// dettagli.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
