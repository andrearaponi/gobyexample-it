// Leggere e scrivere file sono compiti base necessari per
// molti programmi Go. Prima esamineremo alcuni esempi di
// lettura di file.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Leggere file richiede di controllare la maggior parte delle chiamate per errori.
// Questo helper semplificherà i nostri controlli di errore qui sotto.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Forse il compito più basilare di lettura file è
	// inghiottire l'intero contenuto di un file in memoria.
	dat, err := os.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	// Spesso vorrai più controllo su come e quali
	// parti di un file vengono lette. Per questi compiti, inizia
	// aprendo (`Open`) un file per ottenere un valore `os.File`.
	f, err := os.Open("/tmp/dat")
	check(err)

	// Leggi alcuni byte dall'inizio del file.
	// Permetti fino a 5 di essere letti ma nota anche quanti
	// sono stati effettivamente letti.
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// Puoi anche fare `Seek` a una posizione conosciuta nel file
	// e leggere (`Read`) da lì.
	o2, err := f.Seek(6, io.SeekStart)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// Altri metodi di seek sono relativi alla
	// posizione corrente del cursore,
	_, err = f.Seek(2, io.SeekCurrent)
	check(err)

	// e relativi alla fine del file.
	_, err = f.Seek(-4, io.SeekEnd)
	check(err)

	// Il pacchetto `io` fornisce alcune funzioni che possono
	// essere utili per la lettura di file. Per esempio, le letture
	// come quelle sopra possono essere implementate più robustamente
	// con `ReadAtLeast`.
	o3, err := f.Seek(6, io.SeekStart)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// Non c'è un rewind integrato, ma
	// `Seek(0, io.SeekStart)` realizza questo.
	_, err = f.Seek(0, io.SeekStart)
	check(err)

	// Il pacchetto `bufio` implementa un reader
	// bufferizzato che può essere utile sia per la sua efficienza
	// con molte piccole letture che per i metodi di lettura
	// aggiuntivi che fornisce.
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// Chiudi il file quando hai finito (di solito questo verrebbe
	// programmato immediatamente dopo l'apertura (`Open`) con
	// `defer`).
	f.Close()
}
