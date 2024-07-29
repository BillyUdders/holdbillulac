package v1

import (
	"errors"
	"holdbillulac/api/common"
	"html/template"
	"net/http"
)

var (
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
			<td id="delete-button"><button hx-delete="/player/{{.ID}}" hx-target="closest tr" hx-swap="outerHTML">Remove</button></td>
		</tr>
	`))
)

func getPlayer(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		common.HandleError(errLog, w, errors.New("must supply ID"), http.StatusBadRequest)
		return
	}
	player, err := common.Query[Player](db, selectByID, id)
	if err != nil {
		common.HandleError(errLog, w, err, http.StatusInternalServerError)
		return
	}
	err = trTemplate.Execute(w, player)
	if err != nil {
		common.HandleError(errLog, w, err, http.StatusInternalServerError)
		return
	}
	infoLog.Printf("Get all: %v", player)
}

func getPlayers(w http.ResponseWriter, _ *http.Request) {
	players, err := common.Query[[]Player](db, selectAll)
	if err != nil {
		common.HandleError(errLog, w, err, http.StatusInternalServerError)
		return
	}
	for i := range players {
		player := players[i]
		err = trTemplate.Execute(w, player)
		if err != nil {
			common.HandleError(errLog, w, err, http.StatusInternalServerError)
			return
		}
	}
	infoLog.Printf("Players returned: %v", len(players))
}

func createPlayer(w http.ResponseWriter, r *http.Request) {
	player, err := new(Player).fromBody(r.Body)
	if err != nil {
		common.HandleError(errLog, w, err, http.StatusBadRequest)
		return
	}
	insertId, err := common.Insert(db, insert, player)
	if err != nil {
		common.HandleError(errLog, w, err, http.StatusInternalServerError)
		return
	}
	player.ID = insertId
	err = trTemplate.Execute(w, player)
	if err != nil {
		common.HandleError(errLog, w, err, http.StatusInternalServerError)
		return
	}
	infoLog.Printf("Created: %v", player)
}

func deletePlayer(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	_, err := db.Exec(deleteByID, id)
	if err != nil {
		common.HandleError(errLog, w, err, http.StatusInternalServerError)
		return
	}
	infoLog.Printf("Deleted ID: %v", id)
}
