package main

import (
	rice "github.com/GeertJohan/go.rice"
	"html/template"
	"log"
	"net/http"
	"slices"
)

type Row struct {
	Name string
	Age  int
}

var (
	box  *rice.Box
	rows []Row

	dummyRow = Row{
		Name: "Jeffrey Epsteinmanhower",
		Age:  52,
	}
	rowTemplate = template.Must(template.New("row-partial").Parse(`
		<tr>
			<td>{{.Name}}</td>
			<td>{{.Age}}</td>
			<td><button hx-delete="/rows" hx-target="closest tr" hx-swap="outerHTML">Remove</button></td>
		</tr>
	`))
)

func getRows(w http.ResponseWriter, _ *http.Request) {
	for range rows {
		err := rowTemplate.Execute(w, dummyRow)
		if err != nil {
			return
		}
	}
}

func createRow(w http.ResponseWriter, _ *http.Request) {
	rows = append(rows, dummyRow)
	err := rowTemplate.Execute(w, dummyRow)
	if err != nil {
		return
	}
}

func deleteRow(w http.ResponseWriter, _ *http.Request) {
	rows = slices.Delete(rows, 0, 1)
}

func index(w http.ResponseWriter, _ *http.Request) {
	content := box.MustString("index.html")
	tmpl := template.Must(template.New("index").Parse(content))
	data := map[string]string{
		"Title":       "Big Test Page",
		"Description": "Holden x Bill: Aligulac",
	}
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	addr := "localhost:8080"
	box = rice.MustFindBox("templates")

	http.HandleFunc("GET /", index)
	http.HandleFunc("GET /rows", getRows)
	http.HandleFunc("POST /rows", createRow)
	http.HandleFunc("DELETE /rows", deleteRow)

	log.Printf("Listening on: %s", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
