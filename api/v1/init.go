package v1

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/jmoiron/sqlx"
	"holdbillulac/api/common"
	"log"
	"net/http"
)

var (
	db      *sqlx.DB
	box     *rice.Box
	infoLog *log.Logger
	errLog  *log.Logger
)

func Initialize(dbName string, _infoLog *log.Logger, _errLog *log.Logger) {
	db = common.InitDB(dbName, createTableStmt, _infoLog)
	box = rice.MustFindBox("../templates")
	infoLog = _infoLog
	errLog = _errLog

	http.HandleFunc("GET /", Index)
	http.HandleFunc("GET /rows", GetPlayers)
	http.HandleFunc("POST /rows", CreatePlayer)
	http.HandleFunc("GET /rows/{id}", GetPlayer)
	http.HandleFunc("DELETE /rows/{id}", DeletePlayer)
}
