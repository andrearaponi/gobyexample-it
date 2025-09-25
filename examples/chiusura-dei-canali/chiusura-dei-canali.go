// _Chiudere_ un canale indica che non verranno più inviati
// valori su di esso. Questo può essere utile per comunicare
// il completamento ai riceventi del canale.

package main

import "fmt"

// In questo esempio useremo un canale `jobs` per
// comunicare lavoro da fare dalla goroutine `main()`
// a una goroutine worker. Quando non abbiamo più job per
// il worker, chiuderemo il canale `jobs`.
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// Ecco la goroutine worker. Riceve ripetutamente
	// da `jobs` con `j, more := <-jobs`. In questa
	// forma speciale di ricezione a 2 valori, il valore `more`
	// sarà `false` se `jobs` è stato chiuso e tutti
	// i valori nel canale sono già stati ricevuti.
	// Usiamo questo per notificare su `done` quando abbiamo
	// lavorato tutti i nostri job.
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// Questo invia 3 job al worker tramite il canale `jobs`,
	// poi lo chiude.
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// Aspettiamo il worker usando l'approccio di
	// [sincronizzazione](sincronizzazione-con-canali) che
	// abbiamo visto prima.
	<-done

	// Leggere da un canale chiuso ha successo immediatamente,
	// restituendo il valore zero del tipo sottostante.
	// Il secondo valore di ritorno opzionale è `true` se il
	// valore ricevuto è stato consegnato da un'operazione di invio
	// riuscita al canale, o `false` se era un valore zero
	// generato perché il canale è chiuso e vuoto.
	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)
}
