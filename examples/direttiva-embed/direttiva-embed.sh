# Usa questi comandi per eseguire l'esempio.
# (Nota: a causa delle limitazioni del go playground,
# questo esempio può essere eseguito solo sulla tua macchina locale.)
$ mkdir -p folder
$ echo "hello go" > folder/single_file.txt
$ echo "123" > folder/file1.hash
$ echo "456" > folder/file2.hash

$ go run direttiva-embed.go
hello go
hello go
123
456

