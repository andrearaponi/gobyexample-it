# Per sperimentare con il programma dei flag da linea
# di comando è meglio prima compilarlo e poi eseguire
# direttamente il binario risultante.
$ go build flag-linea-comando.go

# Prova il programma compilato fornendogli prima
# valori per tutti i flag.
$ ./flag-linea-comando -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []

# Nota che se ometti i flag assumono automaticamente
# i loro valori di default.
$ ./flag-linea-comando -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []

# Argomenti posizionali finali possono essere forniti
# dopo qualsiasi flag.
$ ./flag-linea-comando -word=opt a1 a2 a3
word: opt
...
tail: [a1 a2 a3]

# Nota che il pacchetto `flag` richiede che tutti i flag
# appaiano prima degli argomenti posizionali (altrimenti
# i flag saranno interpretati come argomenti posizionali).
$ ./flag-linea-comando -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3 -numb=7]

# Usa i flag `-h` o `--help` per ottenere automaticamente
# il testo di aiuto per il programma da linea di comando.
$ ./flag-linea-comando -h
Usage of ./flag-linea-comando:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string

# Se fornisci un flag che non è stato specificato al
# pacchetto `flag`, il programma stamperà un messaggio
# di errore e mostrerà di nuovo il testo di aiuto.
$ ./flag-linea-comando -wat
flag provided but not defined: -wat
Usage of ./flag-linea-comando:
...
