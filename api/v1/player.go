package v1

import (
	"errors"
	"github.com/gorilla/mux"
	"holdbillulac/api/common"
	"net/http"
)

var playerQueries = common.CRUD{
	Insert:    "INSERT INTO players (name, age, MMR) VALUES (:name, :age, :MMR)",
	SelectAll: "SELECT * FROM players",
	Select:    "SELECT * FROM players WHERE id = ?",
	Delete:    "DELETE FROM players WHERE id = ?",
	Update:    "<NOT IMPLEMENTED>",
}

func getPlayer(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		common.HandleError(errLog, w, errors.New("must supply ID"), http.StatusBadRequest)
		return
	}
	err := common.Get[Player](db, w, playerQueries.Select, id, playerTr)
	if err != nil {
		common.HandleError(errLog, w, err, http.StatusInternalServerError)
		return
	}
}

func getPlayers(w http.ResponseWriter, _ *http.Request) {
	err := common.GetAll[Player](db, w, playerQueries.SelectAll, playerTr)
	if err != nil {
		common.HandleError(errLog, w, err, http.StatusInternalServerError)
		return
	}
}

func createPlayer(w http.ResponseWriter, r *http.Request) {
	player, err := new(Player).fromBody(r.Body)
	if err != nil {
		common.HandleError(errLog, w, err, http.StatusBadRequest)
		return
	}
	err = common.Create[Player](db, w, playerQueries.Insert, *player, playerTr)
	if err != nil {
		common.HandleError(errLog, w, err, http.StatusInternalServerError)
		return
	}
}

func deletePlayer(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := db.Exec(playerQueries.Delete, id)
	if err != nil {
		common.HandleError(errLog, w, err, http.StatusInternalServerError)
		return
	}
	infoLog.Printf("Deleted ID: %v", id)
}
