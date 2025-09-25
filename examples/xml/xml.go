// Go offre supporto integrato per XML e formati
// simili a XML con il pacchetto `encoding/xml`.

package main

import (
	"encoding/xml"
	"fmt"
)

// Plant sarà mappato a XML. Similmente agli
// esempi JSON, i tag dei campi contengono direttive per
// encoder e decoder. Qui usiamo alcune funzionalità speciali
// del pacchetto XML: il nome del campo `XMLName` detta
// il nome dell'elemento XML che rappresenta questo struct;
// `id,attr` significa che il campo `Id` è un
// _attributo_ XML piuttosto che un elemento annidato.
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	// Emette XML che rappresenta la nostra pianta; usando
	// `MarshalIndent` per produrre un output
	// più leggibile dall'uomo.
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))

	// Per aggiungere un header XML generico all'output, aggiungilo
	// esplicitamente.
	fmt.Println(xml.Header + string(out))

	// Usa `Unmarshal` per parsare un flusso di byte con XML
	// in una struttura dati. Se l'XML è malformato o
	// non può essere mappato su Plant, sarà restituito
	// un errore descrittivo.
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	// Il tag del campo `parent>child>plant` dice all'encoder
	// di annidare tutti i `plant` sotto `<parent><child>...`
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}
