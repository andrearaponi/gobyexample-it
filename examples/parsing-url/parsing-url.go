// Gli URL forniscono un [modo uniforme per localizzare risorse](https://adam.herokuapp.com/past/2010/3/30/urls_are_the_uniform_way_to_locate_resources/).
// Ecco come parsare URL in Go.

package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	// Parseremo questo URL di esempio, che include
	// schema, info di autenticazione, host, porta, percorso,
	// parametri di query e frammento di query.
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// Parsa l'URL e assicurati che non ci siano errori.
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// Accedere allo schema è semplice.
	fmt.Println(u.Scheme)

	// `User` contiene tutte le info di autenticazione; chiama
	// `Username` e `Password` su questo per i valori
	// individuali.
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// `Host` contiene sia l'hostname che la porta,
	// se presente. Usa `SplitHostPort` per estrarli.
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// Qui estraiamo il `path` e il frammento dopo
	// il `#`.
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// Per ottenere parametri di query in formato stringa `k=v`,
	// usa `RawQuery`. Puoi anche parsare parametri di query
	// in una mappa. Le mappe dei parametri di query parsati vanno da
	// stringhe a slice di stringhe, quindi indicizza in `[0]`
	// se vuoi solo il primo valore.
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
