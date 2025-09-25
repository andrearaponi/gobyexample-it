// In questo esempio vedremo come implementare
// un _pool di worker_ usando goroutine e canali.

package main

import (
	"fmt"
	"time"
)

// Ecco il worker, di cui eseguiremo diverse
// istanze concorrenti. Questi worker riceveranno
// lavoro dal canale `jobs` e invieranno i risultati
// corrispondenti su `results`. Aspetteremo un secondo per job
// per simulare un'operazione costosa.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// Per utilizzare il nostro pool di worker dobbiamo inviare
	// loro del lavoro e raccogliere i risultati. Creiamo 2
	// canali per questo.
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Questo avvia 3 worker, inizialmente bloccati
	// perché non ci sono ancora job.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Qui inviamo 5 `jobs` e poi chiudiamo quel
	// canale per indicare che è tutto il lavoro che abbiamo.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Infine raccogliamo tutti i risultati del lavoro.
	// Questo assicura anche che le goroutine worker abbiano
	// terminato. Un modo alternativo per attendere più
	// goroutine è utilizzare un [WaitGroup](waitgroups).
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
