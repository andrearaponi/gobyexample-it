// Gli [_hash SHA256_](https://en.wikipedia.org/wiki/SHA-2) sono
// frequentemente usati per calcolare identità brevi per blob
// binari o di testo. Per esempio, i certificati TLS/SSL usano SHA256
// per calcolare la firma di un certificato. Ecco come calcolare
// hash SHA256 in Go.

package main

// Go implementa diverse funzioni hash in vari
// pacchetti `crypto/*`.
import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "sha256 this string"

	// Qui iniziamo con un nuovo hash.
	h := sha256.New()

	// `Write` si aspetta byte. Se hai una stringa `s`,
	// usa `[]byte(s)` per convertirla in byte.
	h.Write([]byte(s))

	// Questo ottiene il risultato hash finalizzato come slice
	// di byte. L'argomento di `Sum` può essere usato per aggiungere
	// a una slice di byte esistente: di solito non è necessario.
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
