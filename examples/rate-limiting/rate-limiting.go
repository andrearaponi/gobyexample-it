// Il [_rate limiting_](https://en.wikipedia.org/wiki/Rate_limiting)
// è un meccanismo importante per controllare l'utilizzo delle risorse
// e mantenere la qualità del servizio. Go
// supporta elegantemente il rate limiting con goroutine,
// canali e [ticker](ticker).

package main

import (
	"fmt"
	"time"
)

func main() {

	// Prima vedremo il rate limiting di base. Supponiamo
	// di voler limitare la gestione delle richieste in arrivo.
	// Serviremo queste richieste da un canale con lo
	// stesso nome.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Questo canale `limiter` riceverà un valore
	// ogni 200 millisecondi. Questo è il regolatore nel
	// nostro schema di rate limiting.
	limiter := time.Tick(200 * time.Millisecond)

	// Bloccandoci su una ricezione dal canale `limiter`
	// prima di servire ogni richiesta, ci limitiamo a
	// 1 richiesta ogni 200 millisecondi.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// Potremmo voler consentire brevi raffiche di richieste nel
	// nostro schema di rate limiting preservando il
	// limite generale del rate. Possiamo farlo
	// bufferizzando il nostro canale limiter. Questo canale `burstyLimiter`
	// consentirà raffiche fino a 3 eventi.
	burstyLimiter := make(chan time.Time, 3)

	// Riempiamo il canale per rappresentare la raffica consentita.
	for range 3 {
		burstyLimiter <- time.Now()
	}

	// Ogni 200 millisecondi proveremo ad aggiungere un nuovo
	// valore a `burstyLimiter`, fino al suo limite di 3.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Ora simuliamo altre 5 richieste in arrivo. Le prime
	// 3 di queste beneficeranno della capacità di raffica
	// di `burstyLimiter`.
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
