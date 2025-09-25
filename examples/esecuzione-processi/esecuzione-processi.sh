# Quando eseguiamo il nostro programma viene sostituito da `ls`.
$ go run esecuzione-processi.go
total 16
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 esecuzione-processi.go

# Nota che Go non offre una funzione Unix `fork`
# classica. Di solito questo non è un problema però, dato che
# avviare goroutine, creare processi e eseguire
# processi copre la maggior parte dei casi d'uso per `fork`.
