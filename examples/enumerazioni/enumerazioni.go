// I _tipi enumerati_ (enum) sono un caso speciale di
// [tipi somma](https://en.wikipedia.org/wiki/Algebraic_data_type).
// Un enum è un tipo che ha un numero fisso di possibili
// valori, ognuno con un nome distinto. Go non ha un
// tipo enum come funzionalità linguistica distinta, ma gli enum
// sono semplici da implementare usando idiomi linguistici esistenti.

package main

import "fmt"

// Il nostro tipo enum `ServerState` ha un tipo `int` sottostante.
type ServerState int

// I possibili valori per `ServerState` sono definiti come
// costanti. La parola chiave speciale [iota](https://go.dev/ref/spec#Iota)
// genera valori costanti successivi automaticamente; in questo
// caso 0, 1, 2 e così via.
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

// Implementando l'interfaccia [fmt.Stringer](https://pkg.go.dev/fmt#Stringer),
// i valori di `ServerState` possono essere stampati o convertiti
// in stringhe.
//
// Questo può diventare ingombrante se ci sono molti valori possibili. In tali
// casi lo [strumento stringer](https://pkg.go.dev/golang.org/x/tools/cmd/stringer)
// può essere usato insieme a `go:generate` per automatizzare il
// processo. Vedi [questo post](https://eli.thegreenplace.net/2021/a-comprehensive-guide-to-go-generate)
// per una spiegazione più lunga.
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)
	// Se abbiamo un valore di tipo `int`, non possiamo passarlo a `transition` - il
	// compilatore si lamenterà per la mancata corrispondenza di tipo. Questo fornisce un certo grado di
	// sicurezza dei tipi a tempo di compilazione per gli enum.

	ns2 := transition(ns)
	fmt.Println(ns2)
}

// transition emula una transizione di stato per un
// server; prende lo stato esistente e restituisce
// un nuovo stato.
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		// Supponiamo di controllare alcuni predicati qui per
		// determinare il prossimo stato...
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
