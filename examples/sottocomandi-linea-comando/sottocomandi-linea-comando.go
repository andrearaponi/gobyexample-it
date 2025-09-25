// Alcuni tool da linea di comando, come `go` o `git`
// hanno molti *sottocomandi*, ognuno con il proprio
// set di flag. Per esempio, `go build` e `go get` sono
// due sottocomandi diversi del tool `go`.
// Il pacchetto `flag` ci permette di definire facilmente
// sottocomandi semplici che hanno i propri flag.

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// Dichiariamo un sottocomando usando la funzione
	// `NewFlagSet`, e procediamo a definire nuovi flag
	// specifici per questo sottocomando.
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	// Per un sottocomando diverso possiamo definire
	// flag supportati diversi.
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	// Il sottocomando è atteso come primo argomento
	// del programma.
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	// Controlla quale sottocomando è invocato.
	switch os.Args[1] {

	// Per ogni sottocomando, parsiamo i suoi flag
	// e abbiamo accesso agli argomenti posizionali finali.
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}
