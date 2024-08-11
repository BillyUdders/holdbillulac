package v1

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"net/http"
)

var (
	db  *sqlx.DB
	box *rice.Box
)

func Initialize(_db *sqlx.DB) *mux.Router {
	db = _db
	box = rice.MustFindBox("../static")

	r := mux.NewRouter()

	// Serve static files from the /static/ directory.
	cssFileServer := http.StripPrefix("/static/", http.FileServer(box.HTTPBox()))
	r.PathPrefix("/static/").Handler(cssFileServer)

	// Serve index template
	r.HandleFunc("/", index).Methods("GET")

	playerRoutes := r.PathPrefix("/player").Subrouter()
	playerRoutes.HandleFunc("", getPlayers).Methods("GET")
	playerRoutes.HandleFunc("", createPlayer).Methods("POST")
	playerRoutes.HandleFunc("/{id}", getPlayer).Methods("GET")
	playerRoutes.HandleFunc("/{id}", deletePlayer).Methods("DELETE")

	navRoutes := r.PathPrefix("/nav").Subrouter()
	navRoutes.HandleFunc("", getNavs).Methods("GET")
	navRoutes.HandleFunc("/{id}", getNav).Methods("GET")

	return r
}
