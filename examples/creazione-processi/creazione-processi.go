// A volte i nostri programmi Go devono creare altri
// processi non-Go.

package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {

	// Inizieremo con un comando semplice che non prende
	// argomenti o input e stampa solo qualcosa su
	// stdout. L'helper `exec.Command` crea un oggetto
	// per rappresentare questo processo esterno.
	dateCmd := exec.Command("date")

	// Il metodo `Output` esegue il comando, aspetta che
	// finisca e raccoglie il suo output standard.
	// Se non ci sono errori, `dateOut` conterrà i byte
	// con le informazioni della data.
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// `Output` e altri metodi di `Command` restituiranno
	// `*exec.Error` se c'è stato un problema nell'eseguire il
	// comando (es. percorso sbagliato), e `*exec.ExitError`
	// se il comando è stato eseguito ma è uscito con un codice
	// di ritorno diverso da zero.
	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("failed executing:", err)
		case *exec.ExitError:
			fmt.Println("command exit rc =", e.ExitCode())
		default:
			panic(err)
		}
	}

	// Ora vedremo un caso leggermente più complesso
	// dove inviamo dati al processo esterno sul suo
	// `stdin` e raccogliamo i risultati dal suo `stdout`.
	grepCmd := exec.Command("grep", "hello")

	// Qui prendiamo esplicitamente le pipe di input/output,
	// avviamo il processo, scriviamo dell'input ad esso,
	// leggiamo l'output risultante e infine aspettiamo
	// che il processo esca.
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()

	// Abbiamo omesso i controlli degli errori nell'esempio sopra,
	// ma potresti usare il solito pattern `if err != nil` per
	// tutti. Raccogliamo anche solo i risultati di `StdoutPipe`,
	// ma potresti raccogliere `StderrPipe` nello
	// stesso identico modo.
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// Nota che quando creiamo comandi dobbiamo
	// fornire un comando e un array di argomenti
	// esplicitamente delineati, invece di poter passare
	// una sola stringa da riga di comando. Se vuoi creare un
	// comando completo con una stringa, puoi usare l'opzione
	// `-c` di `bash`:
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
