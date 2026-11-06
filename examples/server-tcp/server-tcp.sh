# Esegui il server TCP in background.
$ go run server-tcp.go &

# Invia dei dati e cattura la risposta usando netcat.
$ echo "Hello from netcat" | nc localhost 8090
ACK: HELLO FROM NETCAT

