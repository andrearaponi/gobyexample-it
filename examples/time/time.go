// Go offre supporto esteso per date/orari e durate;
// ecco alcuni esempi.

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// Inizieremo ottenendo l'ora corrente.
	now := time.Now()
	p(now)

	// Puoi costruire uno struct `time` fornendo
	// anno, mese, giorno, ecc. Le date/orari sono sempre associate
	// a una `Location`, cioè fuso orario.
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// Puoi estrarre i vari componenti del valore
	// data/ora come aspettato.
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// È disponibile anche il `Weekday` lunedì-domenica.
	p(then.Weekday())

	// Questi metodi confrontano due date/orari, testando se il
	// primo avviene prima, dopo, o nello stesso momento
	// del secondo, rispettivamente.
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// Il metodo `Sub` restituisce una `Duration` che rappresenta
	// l'intervallo tra due date/orari.
	diff := now.Sub(then)
	p(diff)

	// Possiamo calcolare la durata della duration in
	// varie unità.
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// Puoi usare `Add` per avanzare una data/ora di una data
	// duration, o con un `-` per muoverti indietro di una
	// duration.
	p(then.Add(diff))
	p(then.Add(-diff))
}
