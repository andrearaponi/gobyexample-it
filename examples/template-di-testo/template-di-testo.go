// Go offre supporto integrato per creare contenuto dinamico o mostrare output personalizzato
// all'utente con il pacchetto `text/template`. Un pacchetto fratello
// chiamato `html/template` fornisce la stessa API ma ha funzionalità di sicurezza aggiuntive
// e dovrebbe essere usato per generare HTML.

package main

import (
	"os"
	"text/template"
)

func main() {

	// Possiamo creare un nuovo template e analizzare il suo corpo da
	// una stringa.
	// I template sono un mix di testo statico e "azioni" racchiuse in
	// `{{...}}` che vengono usate per inserire dinamicamente il contenuto.
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	// In alternativa, possiamo usare la funzione `template.Must` per
	// fare panic nel caso `Parse` restituisca un errore. Questo è specialmente
	// utile per template inizializzati nell'ambito globale.
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	// "Eseguendo" il template generiamo il suo testo con
	// valori specifici per le sue azioni. L'azione `{{.}}` è
	// sostituita dal valore passato come parametro a `Execute`.
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	// Funzione helper che useremo qui sotto.
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	// Se i dati sono una struct possiamo usare l'azione `{{.FieldName}}` per accedere
	// ai suoi campi. I campi dovrebbero essere esportati per essere accessibili quando un
	// template è in esecuzione.
	t2 := Create("t2", "Name: {{.Name}}\n")

	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	// Lo stesso vale per le mappe; con le mappe non c'è restrizione sul
	// case dei nomi delle chiavi.
	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})

	// if/else forniscono esecuzione condizionale per i template. Un valore è considerato
	// false se è il valore di default di un tipo, come 0, una stringa vuota,
	// un puntatore nil, ecc.
	// Questo esempio dimostra un'altra
	// funzionalità dei template: usare `-` nelle azioni per rimuovere gli spazi bianchi.
	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	// I blocchi range ci permettono di iterare attraverso slice, array, mappe o canali. All'interno
	// del blocco range `{{.}}` è impostato sull'elemento corrente dell'iterazione.
	t4 := Create("t4",
		"Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout,
		[]string{
			"Go",
			"Rust",
			"C++",
			"C#",
		})
}
