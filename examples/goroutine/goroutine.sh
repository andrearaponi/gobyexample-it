# Quando eseguiamo questo programma, vediamo prima l'output della
# chiamata bloccante, poi l'output delle due
# goroutine. L'output delle goroutine può essere intercalato,
# perché le goroutine vengono eseguite concorrentemente dal
# runtime di Go.
$ go run goroutine.go
direct : 0
direct : 1
direct : 2
goroutine : 0
going
goroutine : 1
goroutine : 2
done

# Ora vedremo un complemento alle goroutine nei
# programmi Go concorrenti: i canali.
