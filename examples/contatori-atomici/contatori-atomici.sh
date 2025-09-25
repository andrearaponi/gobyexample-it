# Ci aspettiamo di ottenere esattamente 50.000 operazioni. Se avessimo
# usato un integer non atomico e lo avessimo incrementato con
# `ops++`, probabilmente otterremmo un numero diverso,
# che cambia tra le esecuzioni, perché le goroutine
# interferirebbe tra loro. Inoltre, otterremmo
# errori di data race quando eseguiamo con il
# flag `-race`.
$ go run contatori-atomici.go
ops: 50000

# Successivamente vedremo i mutex, un altro strumento per gestire
# lo stato.
