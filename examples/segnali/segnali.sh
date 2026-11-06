# Quando eseguiamo questo programma si bloccherà aspettando un
# segnale. Digitando `ctrl-C` (che il
# terminale mostra come `^C`) possiamo inviare un segnale `SIGINT`,
# facendo stampare al programma la causa della cancellazione
# e poi uscire.
$ go run segnali.go
awaiting signal
^C
interrupt signal received
exiting
