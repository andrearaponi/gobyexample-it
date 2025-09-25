// Un `panic` tipicamente significa che qualcosa è andato inaspettatamente
// storto. Lo usiamo principalmente per fallire velocemente su errori che
// non dovrebbero verificarsi durante l'operazione normale, o che non
// siamo preparati a gestire con grazia.

package main

import "os"

func main() {

	// Useremo panic in tutto questo sito per verificare
	// errori inaspettati. Questo è l'unico programma nel
	// sito progettato per fare panic.
	panic("a problem")

	// Un uso comune di panic è interrompere se una funzione
	// restituisce un valore di errore che non sappiamo come
	// (o vogliamo) gestire. Ecco un esempio di
	// `panic` se otteniamo un errore inaspettato durante la creazione di un nuovo file.
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
