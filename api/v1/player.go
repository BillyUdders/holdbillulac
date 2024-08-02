package v1

import (
	"context"
	"errors"
	"github.com/gorilla/mux"
	"holdbillulac/api/common"
	"net/http"
)

var playerCRUD = common.CRUD{
	Insert:    "INSERT INTO players (name, age, MMR) VALUES (:name, :age, :MMR)",
	SelectAll: "SELECT * FROM players",
	Select:    "SELECT * FROM players WHERE id = ?",
	Delete:    "DELETE FROM players WHERE id = ?",
	Update:    "UPDATE",
}

func getPlayer(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		common.HandleError(errLog, w, errors.New("must supply ID"), http.StatusBadRequest)
		return
	}
	player, err := common.Query[Player](db, playerCRUD.Select, id)
	if err != nil {
		common.HandleError(errLog, w, err, http.StatusInternalServerError)
		return
	}
	err = playerTr(player).Render(context.Background(), w)
	if err != nil {
		common.HandleError(errLog, w, err, http.StatusInternalServerError)
		return
	}
	infoLog.Printf("Get all: %v", player)
}

func getPlayers(w http.ResponseWriter, _ *http.Request) {
	players, err := common.Query[[]Player](db, playerCRUD.SelectAll)
	if err != nil {
		common.HandleError(errLog, w, err, http.StatusInternalServerError)
		return
	}
	for i := range players {
		player := players[i]
		err = playerTr(player).Render(context.Background(), w)
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
	insertId, err := common.Insert(db, playerCRUD.Insert, player)
	if err != nil {
		common.HandleError(errLog, w, err, http.StatusInternalServerError)
		return
	}
	player.ID = insertId
	err = playerTr(*player).Render(context.Background(), w)
	if err != nil {
		common.HandleError(errLog, w, err, http.StatusInternalServerError)
		return
	}
	infoLog.Printf("Created: %v", player)
}

func deletePlayer(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := db.Exec(playerCRUD.Delete, id)
	if err != nil {
		common.HandleError(errLog, w, err, http.StatusInternalServerError)
		return
	}
	infoLog.Printf("Deleted ID: %v", id)
}
