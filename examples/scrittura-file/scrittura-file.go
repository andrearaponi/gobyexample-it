// Scrivere file in Go segue pattern simili a quelli
// che abbiamo visto prima per la lettura.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Per iniziare, ecco come scaricare una stringa (o semplicemente
	// byte) in un file.
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// Per scritture più granulari, apri un file per la scrittura.
	f, err := os.Create("/tmp/dat2")
	check(err)

	// È idiomatico deferire una `Close` immediatamente
	// dopo aver aperto un file.
	defer f.Close()

	// Puoi scrivere (`Write`) slice di byte come ti aspetteresti.
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// È disponibile anche `WriteString`.
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// Esegui un `Sync` per scaricare le scritture su storage stabile.
	f.Sync()

	// `bufio` fornisce writer bufferizzati oltre
	// ai reader bufferizzati che abbiamo visto prima.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	// Usa `Flush` per assicurarti che tutte le operazioni bufferizzate siano
	// state applicate al writer sottostante.
	w.Flush()

}
