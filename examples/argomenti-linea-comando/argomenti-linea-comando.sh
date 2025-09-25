# Per sperimentare con gli argomenti da linea di comando è meglio
# costruire prima un binario con `go build`.
$ go build argomenti-linea-comando.go
$ ./argomenti-linea-comando a b c d
[./argomenti-linea-comando a b c d]       
[a b c d]
c

# Successivamente vedremo elaborazioni più avanzate da linea di comando
# con i flag.
