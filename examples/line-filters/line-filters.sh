# Per provare il nostro line filter, prima crea un file con alcune
# linee minuscole.
$ echo 'hello'   > /tmp/lines
$ echo 'filter' >> /tmp/lines

# Poi usa il line filter per ottenere linee maiuscole.
$ cat /tmp/lines | go run line-filters.go
HELLO
FILTER
