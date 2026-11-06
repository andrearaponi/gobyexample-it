# `zeroval` non modifica la `i` in `main`, mentre
# `zeroptr` sì, perché ha un riferimento
# all'indirizzo di memoria di quella variabile.
$ go run puntatori.go
initial: 1
zeroval: 1
zeroptr: 0
pointer: 0x42131100
value at *p: 42
value at *p: 0