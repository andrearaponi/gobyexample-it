# I programmi creati restituiscono output che è lo stesso
# come se li avessimo eseguiti direttamente dalla riga di comando.
$ go run creazione-processi.go 
> date
Thu 05 May 2022 10:10:12 PM PDT

# date non ha un flag `-x` quindi uscirà con
# un messaggio di errore e codice di ritorno diverso da zero.
command exited with rc = 1
> grep hello
hello grep

> ls -a -l -h
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 creazione-processi.go
