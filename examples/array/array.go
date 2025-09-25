// In Go, un _array_ è una sequenza numerata di elementi di una
// lunghezza specifica. Nel codice Go tipico, le [slice](slices) sono
// molto più comuni; gli array sono utili in alcuni scenari
// speciali.

package main

import "fmt"

func main() {

	// Qui creiamo un array `a` che conterrà esattamente
	// 5 `int`. Il tipo degli elementi e la lunghezza fanno entrambi
	// parte del tipo dell'array. Di default un array ha
	// valore zero, che per gli `int` significa `0`.
	var a [5]int
	fmt.Println("emp:", a)

	// Possiamo impostare un valore a un indice usando la
	// sintassi `array[indice] = valore`, e ottenere un valore con
	// `array[indice]`.
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// La funzione builtin `len` restituisce la lunghezza di un array.
	fmt.Println("len:", len(a))

	// Usa questa sintassi per dichiarare e inizializzare un array
	// in una sola riga.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Puoi anche far contare al compilatore il numero di
	// elementi per te con `...`
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Se specifichi l'indice con `:`, gli elementi
	// intermedi saranno azzerati.
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	// I tipi array sono unidimensionali, ma puoi
	// comporre i tipi per costruire strutture dati
	// multidimensionali.
	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// Puoi anche creare e inizializzare array
	// multidimensionali in una sola volta.
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}
