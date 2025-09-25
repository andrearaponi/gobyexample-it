# Riceviamo i valori `"one"` e poi `"two"` come
# previsto.
$ time go run select.go 
received one
received two

# Nota che il tempo totale di esecuzione è solo ~2 secondi
# poiché entrambi i `Sleep` di 1 e 2 secondi vengono eseguiti
# concorrentemente.
real	0m2.245s
