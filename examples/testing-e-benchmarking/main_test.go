// Il testing unitario è una parte importante dello scrivere
// programmi Go ben strutturati. Il pacchetto `testing`
// fornisce gli strumenti necessari per scrivere unit test
// e il comando `go test` esegue i test.

// Per scopo dimostrativo, questo codice è nel pacchetto
// `main`, ma potrebbe essere qualsiasi pacchetto. Il codice di testing
// tipicamente risiede nello stesso pacchetto del codice che testa.
package main

import (
	"fmt"
	"testing"
)

// Testeremo questa semplice implementazione di un
// minimo tra interi. Tipicamente, il codice che stiamo testando
// sarebbe in un file sorgente chiamato qualcosa come
// `intutils.go`, e il file di test corrispondente sarebbe
// chiamato `intutils_test.go`.
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Un test si crea scrivendo una funzione con un nome
// che inizia con `Test`.
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		// `t.Error*` riporterà i fallimenti dei test ma continuerà
		// ad eseguire il test. `t.Fatal*` riporterà i fallimenti
		// dei test e fermerà il test immediatamente.
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// Scrivere test può essere ripetitivo, quindi è idiomatico usare
// uno *stile table-driven*, dove gli input dei test e
// gli output attesi sono elencati in una tabella e un singolo ciclo
// li attraversa ed esegue la logica di test.
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		// `t.Run` abilita l'esecuzione di "subtest", uno per ogni
		// entry della tabella. Questi sono mostrati separatamente
		// quando si esegue `go test -v`.
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// I test di benchmark tipicamente vanno nei file `_test.go` e sono
// nominati iniziando con `Benchmark`.
// Qualsiasi codice necessario per eseguire il benchmark ma che non
// dovrebbe essere misurato va prima di questo ciclo.
func BenchmarkIntMin(b *testing.B) {
	for b.Loop() {
		// Il runner di benchmark eseguirà automaticamente questo corpo
		// di ciclo molte volte per determinare una stima ragionevole del
		// tempo di esecuzione di una singola iterazione.
		IntMin(1, 2)
	}
}
