// Il parsing di numeri da stringhe è un compito base ma comune
// in molti programmi; ecco come farlo in Go.

package main

// Il pacchetto built-in `strconv` fornisce il parsing
// dei numeri.
import (
	"fmt"
	"strconv"
)

func main() {

	// Con `ParseFloat`, questo `64` indica quanti bit di
	// precisione parsare.
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// Per `ParseInt`, lo `0` significa dedurre la base dalla
	// stringa. `64` richiede che il risultato stia in 64
	// bit.
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// `ParseInt` riconoscerà i numeri in formato esadecimale.
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// È disponibile anche `ParseUint`.
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// `Atoi` è una funzione di convenienza per il parsing
	// base di `int` in base 10.
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// Le funzioni di parse restituiscono un errore su input non valido.
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
