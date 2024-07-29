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
