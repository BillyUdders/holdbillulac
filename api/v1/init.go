package v1

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

var (
	db      *sqlx.DB
	box     *rice.Box
	infoLog *log.Logger
	errLog  *log.Logger
)

func Initialize(_db *sqlx.DB, _infoLog *log.Logger, _errLog *log.Logger) {
	db = _db
	box = rice.MustFindBox("../templates")
	infoLog = _infoLog
	errLog = _errLog

	http.HandleFunc("GET /", index)
	http.HandleFunc("GET /rows", getPlayers)
	http.HandleFunc("POST /rows", createPlayer)
	http.HandleFunc("GET /rows/{id}", getPlayer)
	http.HandleFunc("DELETE /rows/{id}", deletePlayer)
}
