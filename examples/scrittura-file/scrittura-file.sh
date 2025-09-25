# Prova a eseguire il codice di scrittura file.
$ go run scrittura-file.go 
wrote 5 bytes
wrote 7 bytes
wrote 9 bytes

# Poi controlla il contenuto dei file scritti.
$ cat /tmp/dat1
hello
go
$ cat /tmp/dat2
some
writes
buffered

# Prossimamente esamineremo l'applicazione di alcune delle idee di I/O file
# che abbiamo appena visto agli stream `stdin` e `stdout`.
