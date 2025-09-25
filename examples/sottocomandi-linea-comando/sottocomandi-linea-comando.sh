$ go build sottocomandi-linea-comando.go 

# Prima invoca il sottocomando foo.
$ ./sottocomandi-linea-comando foo -enable -name=joe a1 a2
subcommand 'foo'
  enable: true
  name: joe
  tail: [a1 a2]

# Ora prova bar.
$ ./sottocomandi-linea-comando bar -level 8 a1
subcommand 'bar'
  level: 8
  tail: [a1]

# Ma bar non accetterà i flag di foo.
$ ./sottocomandi-linea-comando bar -enable a1
flag provided but not defined: -enable
Usage of bar:
  -level int
    	level

# Successivamente vedremo le variabili d'ambiente, un altro
# modo comune per parametrizzare i programmi.
