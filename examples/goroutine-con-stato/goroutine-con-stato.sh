# Eseguendo il nostro programma vediamo che l'esempio
# di gestione dello stato basato su goroutine completa circa 80.000
# operazioni totali.
$ go run goroutine-con-stato.go
readOps: 71708
writeOps: 7177

# Per questo caso particolare l'approccio basato su goroutine
# è stato un po' più complesso di quello basato su mutex. Potrebbe
# essere utile in certi casi però, per esempio
# dove hai altri canali coinvolti o quando gestire
# più mutex del genere sarebbe soggetto a errori. Dovresti
# usare l'approccio che ti sembra più naturale, specialmente
# riguardo alla comprensione della correttezza del tuo
# programma.
