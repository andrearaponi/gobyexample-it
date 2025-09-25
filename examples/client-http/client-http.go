// La libreria standard di Go fornisce un ottimo supporto
// per client e server HTTP nel pacchetto `net/http`.
// In questo esempio lo useremo per effettuare semplici
// richieste HTTP.
package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	// Effettua una richiesta HTTP GET a un server. `http.Get` è
	// una scorciatoia comoda che evita di creare un oggetto
	// `http.Client` e chiamare il suo metodo `Get`; usa
	// l'oggetto `http.DefaultClient` che ha impostazioni
	// predefinite utili.
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Stampa lo status della risposta HTTP.
	fmt.Println("Response status:", resp.Status)

	// Stampa le prime 5 righe del body della risposta.
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
