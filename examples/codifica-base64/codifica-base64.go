// Go fornisce supporto integrato per la [codifica/decodifica
// base64](https://en.wikipedia.org/wiki/Base64).

package main

// Questa sintassi importa il pacchetto `encoding/base64` con
// il nome `b64` invece del default `base64`. Ci farà
// risparmiare spazio qui sotto.
import (
	b64 "encoding/base64"
	"fmt"
)

func main() {

	// Ecco la `string` che codificheremo/decodificheremo.
	data := "abc123!?$*&()'-=@~"

	// Go supporta sia base64 standard che compatibile con URL.
	// Ecco come codificare usando l'encoder standard.
	// L'encoder richiede un `[]byte` quindi convertiamo
	// la nostra `string` a quel tipo.
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// La decodifica può restituire un errore, che puoi controllare
	// se non sai già che l'input sia
	// ben formato.
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// Questo codifica/decodifica usando un formato base64
	// compatibile con URL.
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
