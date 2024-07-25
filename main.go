package main

import (
	rice "github.com/GeertJohan/go.rice"
	"html/template"
	"log"
	"net/http"
)

var box *rice.Box
var counter = 1

func createRow(w http.ResponseWriter, _ *http.Request) {
	counter++
	log.Printf("Counter incremented to: %d", counter)
	ret := `
		<tr>
			<td>Rhys Davies</td>
			<td>52223423</td>
			<td><button hx-delete="/rows" hx-target="closest tr" hx-swap="outerHTML">Remove</button></td>
		</tr>
	`
	_, _ = w.Write([]byte(ret))

}

func deleteRow(w http.ResponseWriter, _ *http.Request) {
	counter--
	log.Printf("Counter decremented to: %d", counter)
}

func getAll(w http.ResponseWriter, _ *http.Request) {
	var ret string
	for i := 0; i < counter; i++ {
		ret += `
			<tr>
				<td>Rhys Davies</td>
				<td>52223423</td>
				<td><button hx-delete="/rows" hx-target="closest tr" hx-swap="outerHTML">Remove</button></td>
			</tr>
		`
	}
	log.Printf("Returning: %s", ret)
	_, _ = w.Write([]byte(ret))
}

func index(w http.ResponseWriter, _ *http.Request) {
	content := box.MustString("index.html")
	tmpl := template.Must(template.New("index").Parse(content))
	data := map[string]string{
		"Title":       "Big Test Page",
		"Description": "Rhys Testing out HTMX Golang",
	}
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	box = rice.MustFindBox("templates")
	http.HandleFunc("GET /", index)
	http.HandleFunc("GET /rows", getAll)
	http.HandleFunc("POST /rows", createRow)
	http.HandleFunc("DELETE /rows", deleteRow)

	log.Printf("Listening on: %s", "localhost:8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
