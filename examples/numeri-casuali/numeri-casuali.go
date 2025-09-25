// Il pacchetto `math/rand/v2` di Go fornisce
// generazione di [numeri pseudocasuali](https://en.wikipedia.org/wiki/Pseudorandom_number_generator).

package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	// Per esempio, `rand.IntN` restituisce un `int` casuale n,
	// `0 <= n < 100`.
	fmt.Print(rand.IntN(100), ",")
	fmt.Print(rand.IntN(100))
	fmt.Println()

	// `rand.Float64` restituisce un `float64` `f`,
	// `0.0 <= f < 1.0`.
	fmt.Println(rand.Float64())

	// Questo può essere usato per generare float casuali in
	// altri intervalli, per esempio `5.0 <= f' < 10.0`.
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// Se vuoi un seed conosciuto, crea una nuova
	// `rand.Source` e passala al costruttore `New`.
	// `NewPCG` crea una nuova source
	// [PCG](https://en.wikipedia.org/wiki/Permuted_congruential_generator)
	// che richiede un seed di due numeri `uint64`.
	s2 := rand.NewPCG(42, 1024)
	r2 := rand.New(s2)
	fmt.Print(r2.IntN(100), ",")
	fmt.Print(r2.IntN(100))
	fmt.Println()

	s3 := rand.NewPCG(42, 1024)
	r3 := rand.New(s3)
	fmt.Print(r3.IntN(100), ",")
	fmt.Print(r3.IntN(100))
	fmt.Println()
}
