# Nota che mentre le slice sono tipi diversi dagli array,
# vengono visualizzate in modo simile da `fmt.Println`.
$ go run slice.go
uninit: [] true true
emp: [  ] len: 3 cap: 3
set: [a b c]
get: c
len: 3
apd: [a b c d e f]
cpy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h i]
t == t2
2d:  [[0] [1 2] [2 3 4]]

# Dai un'occhiata a questo [ottimo post del blog](https://go.dev/blog/slices-intro)
# del team di Go per maggiori dettagli sul design e
# l'implementazione delle slice in Go.

# Ora che abbiamo visto array e slice, esamineremo
# l'altra struttura dati builtin fondamentale di Go: le mappe.
