package v1

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"holdbillulac/api/common"
	"io"
	"net/http"
)

// Player type and parsing methods
type Player struct {
	common.Base
	Name string `db:"name"`
	Age  int    `db:"age"`
	MMR  int    `db:"MMR"`
}

func (player *Player) fromBody(body io.ReadCloser) (*Player, error) {
	err := json.NewDecoder(body).Decode(&player)
	if err != nil {
		return nil, err
	}
	if player.Name == "" || player.Age == 0 {
		return nil, errors.New("missing required fields")
	}
	return player, nil
}

func (player *Player) UnmarshalJSON(data []byte) error {
	var raw map[string]string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	player.Name = raw["name"]
	age, err := common.FieldToInt(raw["age"])
	if err != nil {
		return err
	}
	player.Age = age
	mmr, err := common.FieldToInt(raw["mmr"])
	if err != nil {
		return err
	}
	player.MMR = mmr
	return nil
}

// API Queries
var playerQueries = common.CRUD{
	Insert:    "INSERT INTO player (name, age, MMR) VALUES (:name, :age, :MMR)",
	SelectAll: "SELECT * FROM player",
	Select:    "SELECT * FROM player WHERE id = ?",
	Delete:    "DELETE FROM player WHERE id = ?",
	Update:    "<NOT IMPLEMENTED>",
}

// API
func getPlayer(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		common.HandleError(errLog, w, errors.New("must supply ID"), http.StatusBadRequest)
		return
	}
	err := common.Get[*Player](db, w, playerQueries.Select, id, playerTr)
	if err != nil {
		common.HandleError(errLog, w, err, http.StatusInternalServerError)
		return
	}
}

func getPlayers(w http.ResponseWriter, _ *http.Request) {
	err := common.GetAll[*Player](db, w, playerQueries.SelectAll, playerTr)
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
	err = common.Create[*Player](db, w, playerQueries.Insert, player, playerTr)
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
