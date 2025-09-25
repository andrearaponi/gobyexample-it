// Go ha diverse funzioni utili per lavorare con le
// *directory* nel file system.

package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Crea una nuova sotto-directory nella directory
	// di lavoro corrente.
	err := os.Mkdir("subdir", 0755)
	check(err)

	// Quando si creano directory temporanee, è buona
	// pratica `defer` la loro rimozione. `os.RemoveAll`
	// eliminerà un intero albero di directory (simile a
	// `rm -rf`).
	defer os.RemoveAll("subdir")

	// Funzione di aiuto per creare un nuovo file vuoto.
	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	// Possiamo creare una gerarchia di directory, inclusi
	// i genitori con `MkdirAll`. Questo è simile al
	// comando `mkdir -p` da linea di comando.
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// `ReadDir` elenca il contenuto della directory, restituendo uno
	// slice di oggetti `os.DirEntry`.
	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// `Chdir` ci permette di cambiare la directory di lavoro corrente,
	// simile a `cd`.
	err = os.Chdir("subdir/parent/child")
	check(err)

	// Ora vedremo il contenuto di `subdir/parent/child`
	// quando elenchiamo la directory *corrente*.
	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// Torniamo indietro (`cd`) da dove abbiamo iniziato.
	err = os.Chdir("../../..")
	check(err)

	// Possiamo anche visitare una directory *ricorsivamente*,
	// includendo tutte le sue sotto-directory. `WalkDir` accetta
	// una funzione di callback per gestire ogni file o
	// directory visitata.
	fmt.Println("Visiting subdir")
	err = filepath.WalkDir("subdir", visit)
}

// `visit` viene chiamata per ogni file o directory trovata
// ricorsivamente da `filepath.WalkDir`.
func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", path, d.IsDir())
	return nil
}
