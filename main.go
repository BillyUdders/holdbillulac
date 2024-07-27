package main

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

var (
	box     *rice.Box
	db      *sqlx.DB
	errLog  = log.New(log.Writer(), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog = log.New(log.Writer(), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	indexCtx = map[string]string{
		"Title":       "Big Test Page",
		"Description": "Holden x Bill: Aligulac",
	}
)

func main() {
	addr := "localhost:8080"
	box = rice.MustFindBox("templates")
	db = InitDB("rows.db")

	_, err := db.Exec(createTableStmt)
	if err != nil {
		errLog.Fatalf("Unable to migrate table statement %s", createTableStmt)
	}

	http.HandleFunc("GET /", index)
	http.HandleFunc("GET /rows", getPlayers)
	http.HandleFunc("POST /rows", createPlayer)
	http.HandleFunc("GET /rows/{id}", getPlayer)
	http.HandleFunc("DELETE /rows/{id}", deletePlayer)

	infoLog.Printf("Listening on: %s", addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		errLog.Fatal(err)
	}
}
