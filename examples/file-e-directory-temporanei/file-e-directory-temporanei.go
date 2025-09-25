// Durante l'esecuzione del programma, spesso vogliamo creare
// dati che non sono necessari dopo l'uscita del programma.
// *File e directory temporanei* sono utili per questo
// scopo poiché non inquinano il file system nel
// tempo.

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Il modo più semplice per creare un file temporaneo è
	// chiamare `os.CreateTemp`. Crea un file *e*
	// lo apre per lettura e scrittura. Forniamo `""`
	// come primo argomento, così `os.CreateTemp`
	// creerà il file nella posizione predefinita per il nostro OS.
	f, err := os.CreateTemp("", "sample")
	check(err)

	// Mostra il nome del file temporaneo. Sui
	// sistemi operativi Unix la directory sarà probabilmente `/tmp`.
	// Il nome del file inizia con il prefisso dato come
	// secondo argomento a `os.CreateTemp` e il resto
	// viene scelto automaticamente per assicurare che chiamate
	// concorrenti creino sempre nomi di file diversi.
	fmt.Println("Temp file name:", f.Name())

	// Puliamo il file dopo aver finito. Il sistema operativo
	// probabilmente pulirà i file temporanei da solo dopo
	// un po' di tempo, ma è buona pratica farlo
	// esplicitamente.
	defer os.Remove(f.Name())

	// Possiamo scrivere alcuni dati nel file.
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	// Se intendiamo scrivere molti file temporanei, potremmo
	// preferire creare una *directory* temporanea.
	// Gli argomenti di `os.MkdirTemp` sono gli stessi di
	// `CreateTemp`, ma restituisce un *nome* di directory
	// piuttosto che un file aperto.
	dname, err := os.MkdirTemp("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	// Ora possiamo sintetizzare nomi di file temporanei
	// prefissandoli con la nostra directory temporanea.
	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}
