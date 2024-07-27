package main

import (
	"encoding/json"
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
	insert     = "INSERT INTO players (name, age, MMR) VALUES (:name, :age, :MMR)"
	selectAll  = "SELECT * FROM players"
	selectByID = "SELECT * FROM players WHERE id = ?"
	deleteByID = "DELETE FROM players WHERE id = ?"

	trTemplate = template.Must(template.New("player-table-row").Parse(`
		<tr>
			<td id="id">{{.ID}}</td>
			<td id="name">{{.Name}}</td>
			<td id="age">{{.Age}}</td>
			<td id="MMR">{{.MMR}}</td>
			<td id="delete-button"><button hx-delete="/rows/{{.ID}}" hx-target="closest tr" hx-swap="outerHTML">Remove</button></td>
		</tr>
	`))
)

type Player struct {
	Base
	Name string `db:"name"`
	Age  string `db:"age"`
	MMR  string `db:"MMR"`
}

func getPlayer(w http.ResponseWriter, r *http.Request) {
	player, err := Query[Player](db, selectByID, r.PathValue("id"))
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	err = trTemplate.Execute(w, player)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	infoLog.Printf("Get all: %v", player)
}

func getPlayers(w http.ResponseWriter, _ *http.Request) {
	players, err := Query[[]Player](db, selectAll)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	for i := range players {
		player := players[i]
		err = trTemplate.Execute(w, player)
		if err != nil {
			handleError(w, err, http.StatusInternalServerError)
			return
		}
	}
	infoLog.Printf("Players returned: %v", len(players))
}

func createPlayer(w http.ResponseWriter, r *http.Request) {
	var player Player
	err := json.NewDecoder(r.Body).Decode(&player)
	if err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}
	insertId, err := Insert(db, insert, player)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	player.ID = insertId
	err = trTemplate.Execute(w, player)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	infoLog.Printf("Created: %v", player)
}

func deletePlayer(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	_, err := db.Exec(deleteByID, id)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	infoLog.Printf("Deleted ID: %v", id)
}
