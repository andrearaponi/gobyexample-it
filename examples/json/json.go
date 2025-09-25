// Go offre supporto integrato per la codifica e
// decodifica JSON, incluso da e verso tipi di dati
// built-in e personalizzati.

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Useremo questi due struct per dimostrare la codifica e
// decodifica di tipi personalizzati qui sotto.
type response1 struct {
	Page   int
	Fruits []string
}

// Solo i campi esportati saranno codificati/decodificati in JSON.
// I campi devono iniziare con lettere maiuscole per essere esportati.
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	// Prima esamineremo la codifica dei tipi di dati base in
	// stringhe JSON. Ecco alcuni esempi per valori
	// atomici.
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// E qui ne abbiamo alcuni per slice e mappe, che codificano
	// in array e oggetti JSON come ti aspetteresti.
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// Il pacchetto JSON può codificare automaticamente i tuoi
	// tipi di dati personalizzati. Includerà solo i campi esportati
	// nell'output codificato e userà per default
	// quei nomi come chiavi JSON.
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// Puoi usare tag nelle dichiarazioni dei campi struct
	// per personalizzare i nomi delle chiavi JSON codificate. Controlla la
	// definizione di `response2` sopra per vedere un esempio
	// di tali tag.
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// Ora esaminiamo la decodifica dei dati JSON in valori
	// Go. Ecco un esempio per una struttura dati
	// generica.
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// Dobbiamo fornire una variabile dove il pacchetto JSON
	// può mettere i dati decodificati. Questa
	// `map[string]interface{}` conterrà una mappa di stringhe
	// a tipi di dati arbitrari.
	var dat map[string]interface{}

	// Ecco la decodifica vera e propria, e un controllo per
	// errori associati.
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// Per usare i valori nella mappa decodificata,
	// dovremo convertirli al loro tipo appropriato.
	// Per esempio qui convertiamo il valore in `num` al
	// tipo `float64` atteso.
	num := dat["num"].(float64)
	fmt.Println(num)

	// Accedere a dati annidati richiede una serie di
	// conversioni.
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// Possiamo anche decodificare JSON in tipi di dati personalizzati.
	// Questo ha i vantaggi di aggiungere ulteriore
	// type-safety ai nostri programmi ed eliminare il
	// bisogno di type assertion quando accediamo ai dati
	// decodificati.
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// Negli esempi sopra abbiamo sempre usato byte e
	// stringhe come intermediari tra i dati e la
	// rappresentazione JSON su standard out. Possiamo anche
	// fare stream di codifiche JSON direttamente a `os.Writer` come
	// `os.Stdout` o anche body di risposte HTTP.
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

	// Le letture in streaming da `os.Reader` come `os.Stdin`
	// o body di richieste HTTP si fanno con `json.Decoder`.
	dec := json.NewDecoder(strings.NewReader(str))
	res1 := response2{}
	dec.Decode(&res1)
	fmt.Println(res1)
}
