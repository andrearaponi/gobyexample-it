# Quando eseguiamo il programma il messaggio `"ping"` viene
# passato con successo da una goroutine a un'altra tramite
# il nostro canale.
$ go run canali.go 
ping

# Per default invii e ricezioni si bloccano finché sia il
# mittente che il ricevente sono pronti. Questa proprietà ci ha
# permesso di aspettare alla fine del nostro programma il messaggio
# `"ping"` senza dover usare altre sincronizzazioni.
