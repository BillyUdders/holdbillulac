package v1

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
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

func Initialize(_db *sqlx.DB, _infoLog *log.Logger, _errLog *log.Logger) *mux.Router {
	db = _db
	box = rice.MustFindBox("../templates")
	infoLog = _infoLog
	errLog = _errLog

	r := mux.NewRouter()

	// Serve static files from the /static/ directory.
	cssFileServer := http.StripPrefix("/static/", http.FileServer(box.HTTPBox()))
	r.PathPrefix("/static/").Handler(cssFileServer)

	// Define your routes.
	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/player", getPlayers).Methods("GET")
	r.HandleFunc("/player", createPlayer).Methods("POST")
	r.HandleFunc("/player/{id}", getPlayer).Methods("GET")
	r.HandleFunc("/player/{id}", deletePlayer).Methods("DELETE")

	return r
}
