// Nell'esempio precedente abbiamo visto come
// [creare processi esterni](creazione-processi). Lo
// facciamo quando abbiamo bisogno di un processo esterno accessibile a
// un processo Go in esecuzione. A volte vogliamo solo
// sostituire completamente il processo Go corrente con un altro
// (forse non-Go). Per farlo useremo l'implementazione di Go
// della classica funzione
// <a href="https://en.wikipedia.org/wiki/Exec_(operating_system)"><code>exec</code></a>.

package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {

	// Per il nostro esempio eseguiremo `ls`. Go richiede un
	// percorso assoluto al binario che vogliamo eseguire, quindi
	// useremo `exec.LookPath` per trovarlo (probabilmente
	// `/bin/ls`).
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// `Exec` richiede argomenti in forma di slice (invece
	// di una grande stringa). Daremo a `ls` alcuni
	// argomenti comuni. Nota che il primo argomento dovrebbe
	// essere il nome del programma.
	args := []string{"ls", "-a", "-l", "-h"}

	// `Exec` ha anche bisogno di un set di [variabili d'ambiente](variabili-ambiente)
	// da usare. Qui forniamo semplicemente il nostro
	// ambiente corrente.
	env := os.Environ()

	// Ecco la vera chiamata a `syscall.Exec`. Se questa chiamata ha
	// successo, l'esecuzione del nostro processo finirà
	// qui e sarà sostituita dal processo `/bin/ls -a -l -h`.
	// Se c'è un errore otterremo un valore
	// di ritorno.
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
