// Go supporta la formattazione e il parsing di date tramite
// layout basati su pattern.

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// Ecco un esempio base di formattazione di una data
	// secondo RFC3339, usando la corrispondente costante
	// di layout.
	t := time.Now()
	p(t.Format(time.RFC3339))

	// Il parsing delle date usa gli stessi valori di layout di `Format`.
	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)

	// `Format` e `Parse` usano layout basati su esempi. Di solito
	// userai una costante da `time` per questi layout, ma
	// puoi anche fornire layout personalizzati. I layout devono usare il
	// data di riferimento `Mon Jan 2 15:04:05 MST 2006` per mostrare il
	// pattern con cui formattare/parsare una data/stringa.
	// La data di esempio deve essere esattamente come mostrato: l'anno 2006,
	// 15 per l'ora, Monday per il giorno della settimana, ecc.
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

	// Per rappresentazioni puramente numeriche puoi anche
	// usare la formattazione standard delle stringhe con i componenti
	// estratti del valore data/ora.
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// `Parse` restituirà un errore su input malformato
	// spiegando il problema di parsing.
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}
