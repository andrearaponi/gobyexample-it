// `//go:embed` è una [direttiva del
// compilatore](https://pkg.go.dev/cmd/compile#hdr-Compiler_Directives) che
// permette ai programmi di includere file e cartelle arbitrari nel binario Go al
// momento della build. Leggi di più sulla direttiva embed
// [qui](https://pkg.go.dev/embed).
package main

// Importa il pacchetto `embed`; se non usi identificatori esportati
// da questo pacchetto, puoi fare un import vuoto con `_ "embed"`.
import (
	"embed"
)

// Le direttive `embed` accettano percorsi relativi alla directory contenente il
// file sorgente Go. Questa direttiva incorpora il contenuto del file nella
// variabile `string` che la segue immediatamente.
//
//go:embed folder/single_file.txt
var fileString string

// Oppure incorpora il contenuto del file in un `[]byte`.
//
//go:embed folder/single_file.txt
var fileByte []byte

// Possiamo anche incorporare file multipli o anche cartelle con wildcards. Questo usa
// una variabile del [tipo embed.FS](https://pkg.go.dev/embed#FS), che
// implementa un semplice file system virtuale.
//
//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {

	// Stampa il contenuto di `single_file.txt`.
	print(fileString)
	print(string(fileByte))

	// Recupera alcuni file dalla cartella incorporata.
	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}
