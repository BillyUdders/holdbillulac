package main

import (
	"encoding/json"
	"fmt"
	rice "github.com/GeertJohan/go.rice"
	"github.com/jmoiron/sqlx"
	"html/template"
	"log"
	"net/http"
)

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
	createTableStmt = `
		CREATE TABLE IF NOT EXISTS rows (
			id  integer primary key,
			name text,
			age integer
		)
	`
)

func Render(w http.ResponseWriter, row Row) error {
	html := `
		<tr>
			<td id="id">{{.ID}}</td>
			<td id="name">{{.Name}}</td>
			<td id="age">{{.Age}}</td>
			<td id="delete-button"><button hx-delete="/rows/{{.ID}}" hx-target="closest tr" hx-swap="outerHTML">Remove</button></td>
		</tr>
	`
	return template.Must(template.New(fmt.Sprintf("%v-row", row.ID)).Parse(html)).Execute(w, row)
}

func getRow(w http.ResponseWriter, _ *http.Request) {
	err := Render(w, Row{})
	if err != nil {
		return
	}
}

func getRows(w http.ResponseWriter, _ *http.Request) {
	rows, err := getAll[Row](db, "SELECT * FROM rows")
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
	}
	for i := range rows {
		row := rows[i]
		err = Render(w, row)
		if err != nil {
			return
		}
	}
}

func createRow(w http.ResponseWriter, r *http.Request) {
	row, err := parseBody[Row](r)
	if err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}
	insertId, err := insert(db, "INSERT INTO rows (name, age) VALUES (?, ?)", row.Name, row.Age)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	row.ID = insertId
	err = Render(w, row)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
}

func deleteRow(w http.ResponseWriter, r *http.Request) {
	_, err := db.Exec("DELETE FROM rows WHERE id = ?", r.PathValue("id"))
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
}

func parseBody[T any](r *http.Request) (T, error) {
	var t T
	err := json.NewDecoder(r.Body).Decode(&t)
	return t, err
}

func handleError(w http.ResponseWriter, err error, errCode int) {
	errLog.Println(err.Error())
	http.Error(w, err.Error(), errCode)
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
	http.HandleFunc("POST /rows", createRow)
	http.HandleFunc("GET /rows/{id}", getRow)
	http.HandleFunc("DELETE /rows/{id}", deleteRow)

	log.Printf("Listening on: %s", addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
