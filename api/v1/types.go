package v1

import (
	"encoding/json"
	"errors"
	"holdbillulac/api/common"
	"io"
)

type Player struct {
	common.Base
	Name string `db:"name"`
	Age  int    `db:"age"`
	MMR  string `db:"MMR"`
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
