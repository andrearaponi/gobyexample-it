# Eseguendo questo programma farà sì che faccia panic, stampi
# un messaggio di errore e tracce delle goroutine, ed esca con
# uno status diverso da zero.

# Quando il primo panic in `main` si scatena, il programma esce
# senza raggiungere il resto del codice. Se vuoi
# vedere il programma provare a creare un file temporaneo, commenta
# il primo panic.
$ go run panic.go
panic: a problem

goroutine 1 [running]:
main.main()
	/.../panic.go:12 +0x47
...
exit status 2

# Nota che a differenza di alcuni linguaggi che usano le eccezioni
# per gestire molti errori, in Go è idiomatico
# usare valori di ritorno che indicano errori ovunque possibile.
