# Eseguendo il programma vediamo che raccogliamo il valore
# di `FOO` che abbiamo impostato nel programma, ma
# `BAR` è vuoto.
$ go run variabili-ambiente.go
FOO: 1
BAR: 

# La lista di chiavi nell'ambiente dipenderà dalla
# tua macchina specifica.
TERM_PROGRAM
PATH
SHELL
...
FOO

# Se impostiamo prima `BAR` nell'ambiente, il programma
# in esecuzione raccoglie quel valore.
$ BAR=2 go run variabili-ambiente.go
FOO: 1
BAR: 2
...
