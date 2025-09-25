# Eseguire il programma calcola l'hash e lo stampa in
# formato esadecimale leggibile.
$ go run hash-sha256.go
sha256 this string
1af1dfa857bf1d8814fe1af8983c18080019922e557f15a8a...


# Puoi calcolare altri hash usando un pattern simile a
# quello mostrato sopra. Per esempio, per calcolare
# hash SHA512 importa `crypto/sha512` e usa
# `sha512.New()`.

# Nota che se hai bisogno di hash crittograficamente sicuri,
# dovresti ricercare attentamente la
# [forza dell'hash](https://en.wikipedia.org/wiki/Cryptographic_hash_function)!
