// Il pacchetto `net` fornisce gli strumenti necessari per
// costruire facilmente server socket TCP.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	// `net.Listen` avvia il server sulla rete indicata
	// (TCP) e sull'indirizzo indicato (porta 8090 su tutte
	// le interfacce).
	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatal("Error listening:", err)
	}

	// Chiudi il listener per liberare la porta
	// quando l'applicazione termina.
	defer listener.Close()

	// Esegui un ciclo infinito per accettare nuove
	// connessioni dai client.
	for {
		// Attendi una connessione.
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting conn:", err)
			continue
		}

		// Qui usiamo una goroutine per gestire la connessione,
		// così il ciclo principale può continuare ad accettare
		// altre connessioni.
		go handleConnection(conn)
	}
}

// `handleConnection` gestisce la connessione di un singolo
// client, leggendo una riga di testo inviata dal client e
// restituendo una risposta.
func handleConnection(conn net.Conn) {
	// Chiudere la connessione libera le risorse quando
	// abbiamo finito di interagire con il client.
	defer conn.Close()

	// Usa `bufio.NewReader` per leggere una riga di dati
	// dal client (terminata da un carattere di a capo).
	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Read error: %v", err)
		return
	}

	// Crea e invia una risposta al client,
	// dimostrando la comunicazione bidirezionale.
	ackMsg := strings.ToUpper(strings.TrimSpace(message))
	response := fmt.Sprintf("ACK: %s\n", ackMsg)
	_, err = conn.Write([]byte(response))
	if err != nil {
		log.Printf("Server write error: %v", err)
	}
}
