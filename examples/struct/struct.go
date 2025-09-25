// Le _struct_ di Go sono collezioni tipizzate di campi.
// Sono utili per raggruppare i dati insieme per formare
// record.

package main

import "fmt"

// Questo tipo struct `person` ha i campi `name` e `age`.
type person struct {
	name string
	age  int
}

// `newPerson` costruisce una nuova struct person con il nome dato.
func newPerson(name string) *person {
	// Go è un linguaggio con garbage collection; puoi tranquillamente
	// restituire un puntatore a una variabile locale - verrà pulita
	// dal garbage collector solo quando non ci saranno
	// riferimenti attivi ad essa.
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	// Questa sintassi crea una nuova struct.
	fmt.Println(person{"Bob", 20})

	// Puoi nominare i campi quando inizializzi una struct.
	fmt.Println(person{name: "Alice", age: 30})

	// I campi omessi avranno valore zero.
	fmt.Println(person{name: "Fred"})

	// Un prefisso `&` produce un puntatore alla struct.
	fmt.Println(&person{name: "Ann", age: 40})

	// È idiomatico incapsulare la creazione di nuove struct in funzioni costruttore
	fmt.Println(newPerson("Jon"))

	// Accedi ai campi della struct con un punto.
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// Puoi anche usare i punti con i puntatori struct - i
	// puntatori vengono dereferenziati automaticamente.
	sp := &s
	fmt.Println(sp.age)

	// Le struct sono mutabili.
	sp.age = 51
	fmt.Println(sp.age)

	// Se un tipo struct è usato solo per un singolo valore, non dobbiamo
	// dargli un nome. Il valore può avere un tipo
	// struct anonimo. Questa tecnica è comunemente usata per
	// [test table-driven](testing-and-benchmarking).
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
