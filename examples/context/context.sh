# Esegui il server in background.
$ go run context.go &

# Simula una richiesta client a `/hello`, premendo
# Ctrl+C poco dopo l'inizio per segnalare
# la cancellazione.
$ curl localhost:8090/hello
server: hello handler started
^C
server: context canceled
server: hello handler ended
