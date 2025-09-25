$ echo "hello" > /tmp/dat
$ echo "go" >>   /tmp/dat
$ go run lettura-file.go
hello
go
5 bytes: hello
2 bytes @ 6: go
2 bytes @ 6: go
5 bytes: hello

# Prossimamente esamineremo la scrittura di file.
