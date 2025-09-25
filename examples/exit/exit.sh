# Se esegui `exit.go` usando `go run`, l'exit
# verrà catturato da `go` e stampato.
$ go run exit.go
exit status 3

# Compilando ed eseguendo un binario puoi vedere
# lo status nel terminale.
$ go build exit.go
$ ./exit
$ echo $?
3

# Nota che il `!` del nostro programma non è mai stato stampato.
