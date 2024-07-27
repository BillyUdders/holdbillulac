package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

var (
	createTableStmt = `
		CREATE TABLE IF NOT EXISTS players (
			id  	integer primary key,
			name 	text,
			age 	integer,
		    MMR 	integer
		)
	`
	insertStmt = "INSERT INTO players (name, age, MMR) VALUES (:name, :age, :MMR)"
	asListItem = func(id int64) *template.Template {
		return template.Must(template.New(fmt.Sprintf("%v-row", id)).Parse(`
		<tr>
			<td id="id">{{.ID}}</td>
			<td id="name">{{.Name}}</td>
			<td id="age">{{.Age}}</td>
			<td id="MMR">{{.MMR}}</td>
			<td id="delete-button"><button hx-delete="/rows/{{.ID}}" hx-target="closest tr" hx-swap="outerHTML">Remove</button></td>
		</tr>
	`))
	}
)

type Player struct {
	Base
	Name string `db:"name"`
	Age  string `db:"age"`
	MMR  string `db:"MMR"`
}

func (row Player) Render(w http.ResponseWriter) error {
	return asListItem(row.ID).Execute(w, row)
}

func getPlayer(w http.ResponseWriter, r *http.Request) {
	row, err := Query[Player](db, "SELECT * FROM players WHERE id = ?", r.PathValue("id"))
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	err = row.Render(w)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
}

func getPlayers(w http.ResponseWriter, _ *http.Request) {
	players, err := Query[[]Player](db, "SELECT * FROM players")
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	for i := range players {
		err = players[i].Render(w)
		if err != nil {
			handleError(w, err, http.StatusInternalServerError)
			return
		}
	}
}

func createPlayer(w http.ResponseWriter, r *http.Request) {
	var player Player
	err := json.NewDecoder(r.Body).Decode(&player)
	if err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}
	insertId, err := Insert(db, insertStmt, player)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	player.ID = insertId
	err = player.Render(w)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
}

func deletePlayer(w http.ResponseWriter, r *http.Request) {
	_, err := db.Exec("DELETE FROM rows WHERE id = ?", r.PathValue("id"))
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
}
