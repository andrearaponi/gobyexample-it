# La stringa codifica a valori leggermente diversi con gli
# encoder base64 standard e URL (`+` finale vs `-`)
# ma entrambi decodificano alla stringa originale come desiderato.
$ go run codifica-base64.go
YWJjMTIzIT8kKiYoKSctPUB+
abc123!?$*&()'-=@~

YWJjMTIzIT8kKiYoKSctPUB-
abc123!?$*&()'-=@~
