// Un requisito comune nei programmi è ottenere il numero
// di secondi, millisecondi, o nanosecondi dall'
// [epoch Unix](https://en.wikipedia.org/wiki/Unix_time).
// Ecco come farlo in Go.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Usa `time.Now` con `Unix`, `UnixMilli` o `UnixNano`
	// per ottenere il tempo trascorso dall'epoch Unix in secondi,
	// millisecondi o nanosecondi, rispettivamente.
	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	// Puoi anche convertire secondi o nanosecondi interi
	// dall'epoch nel `time` corrispondente.
	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}
