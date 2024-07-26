package main

import (
	"encoding/json"
	rice "github.com/GeertJohan/go.rice"
	"github.com/jmoiron/sqlx"
	"html/template"
	"log"
	"net/http"
)

var createTableStmt = `
	CREATE TABLE IF NOT EXISTS rows (
		id  integer primary key,
		name text,
		age integer
	)
`

type Row struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
	Age  string `db:"age"`
}

var (
	box    *rice.Box
	db     *sqlx.DB
	errLog = log.New(log.Writer(), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	indexCtx = map[string]string{
		"Title":       "Big Test Page",
		"Description": "Holden x Bill: Aligulac",
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
	rows, err := getAll[Row](db, "SELECT * FROM rows")
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
	}
	for i := range rows {
		err = rowTemplate.Execute(w, rows[i])
		if err != nil {
			return
		}
	}
}

func createRow(w http.ResponseWriter, r *http.Request) {
	var row Row
	err := json.NewDecoder(r.Body).Decode(&row)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	insertId, err := insert(db, "INSERT INTO rows (name, age) VALUES (?, ?)", row.Name, row.Age)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	row.ID = insertId
	err = rowTemplate.Execute(w, row)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
}

func handleError(w http.ResponseWriter, err error, errCode int) {
	errLog.Println(err.Error())
	http.Error(w, err.Error(), errCode)
}

func getRow(w http.ResponseWriter, _ *http.Request) {
	err := rowTemplate.Execute(w, nil)
	if err != nil {
		return
	}
}

func deleteRow(w http.ResponseWriter, _ *http.Request) {
	//handleError(w, errors.New("unimplemented"), 500)
	//_, err := w.Write([]byte("Holdennello"))
	//if err != nil {
	//	return
	//}
}

func index(w http.ResponseWriter, _ *http.Request) {
	content := box.MustString("index.html")
	tmpl := template.Must(template.New("index").Parse(content))
	err := tmpl.Execute(w, indexCtx)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	addr := "localhost:8080"
	box = rice.MustFindBox("templates")
	db = initDB("rows.db")

	_, err := db.Exec(createTableStmt)
	if err != nil {
		log.Fatalf("Unable to migrate table statement %s", createTableStmt)
	}

	http.HandleFunc("GET /", index)
	http.HandleFunc("GET /rows", getRows)
	http.HandleFunc("GET /rows/{id}", getRow)
	http.HandleFunc("POST /rows", createRow)
	http.HandleFunc("DELETE /rows", deleteRow)

	log.Printf("Listening on: %s", addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
