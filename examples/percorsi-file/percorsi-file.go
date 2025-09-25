// Il pacchetto `filepath` fornisce funzioni per parsare
// e costruire *percorsi di file* in modo portabile
// tra sistemi operativi; `dir/file` su Linux vs.
// `dir\file` su Windows, per esempio.
package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	// `Join` dovrebbe essere usato per costruire percorsi in modo
	// portabile. Prende un numero qualsiasi di argomenti
	// e costruisce un percorso gerarchico da essi.
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// Dovresti sempre usare `Join` invece di
	// concatenare `/` o `\` manualmente. Oltre
	// a fornire portabilità, `Join` normalizzerà anche
	// i percorsi rimuovendo separatori superflui
	// e cambi di directory.
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// `Dir` e `Base` possono essere usati per dividere un percorso nella
	// directory e nel file. In alternativa, `Split`
	// restituirà entrambi nella stessa chiamata.
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	// Possiamo controllare se un percorso è assoluto.
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	// Alcuni nomi di file hanno estensioni dopo un punto. Possiamo
	// separare l'estensione da tali nomi con `Ext`.
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// Per trovare il nome del file con l'estensione rimossa,
	// usa `strings.TrimSuffix`.
	fmt.Println(strings.TrimSuffix(filename, ext))

	// `Rel` trova un percorso relativo tra una *base* e un
	// *target*. Restituisce un errore se il target non può
	// essere reso relativo alla base.
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
