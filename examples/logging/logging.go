// La libreria standard di Go fornisce strumenti
// semplici per l'output dei log dai programmi Go, con
// il pacchetto [log](https://pkg.go.dev/log) per
// output libero e il pacchetto
// [log/slog](https://pkg.go.dev/log/slog) per
// output strutturato.
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"log/slog"
)

func main() {

	// Semplicemente invocando funzioni come `Println`
	// dal pacchetto `log` si usa il logger _standard_,
	// che è già preconfigurato per un output di logging
	// ragionevole su `os.Stderr`. Metodi aggiuntivi come
	// `Fatal*` o `Panic*` usciranno dal programma
	// dopo il logging.
	log.Println("standard logger")

	// I logger possono essere configurati con _flag_
	// per impostare il loro formato di output. Di default,
	// il logger standard ha i flag `log.Ldate` e
	// `log.Ltime` impostati, e questi sono raccolti
	// in `log.LstdFlags`. Possiamo cambiare i suoi flag
	// per emettere il tempo con precisione microsecondo.
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	// Supporta anche l'emissione del nome file e
	// della linea da cui la funzione `log` è chiamata.
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	// Può essere utile creare un logger personalizzato
	// e passarlo in giro. Quando creiamo un nuovo logger,
	// possiamo impostare un _prefisso_ per distinguere
	// il suo output da altri logger.
	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")

	// Possiamo impostare il prefisso
	// sui logger esistenti (incluso quello standard)
	// con il metodo `SetPrefix`.
	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")

	// I logger possono avere target di output personalizzati;
	// qualsiasi `io.Writer` funziona.
	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)

	// Questa chiamata scrive l'output del log in `buf`.
	buflog.Println("hello")

	// Questo lo mostrerà effettivamente sull'output standard.
	fmt.Print("from buflog:", buf.String())

	// Il pacchetto `slog` fornisce output di log
	// _strutturato_. Per esempio, il logging in
	// formato JSON è semplice.
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")

	// Oltre al messaggio, l'output di `slog` può
	// contenere un numero arbitrario di coppie
	// chiave=valore.
	myslog.Info("hello again", "key", "val", "age", 25)
}
