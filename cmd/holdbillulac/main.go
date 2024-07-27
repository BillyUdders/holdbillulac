package main

import (
	"embed"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	"holdbillulac/api/common"
	api "holdbillulac/api/v1"
	"log"
	"net/http"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func runGooseMigration(db *sqlx.DB) {
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("sqlite3"); err != nil {
		panic(err)
	}
	if err := goose.Up(db.DB, "migrations"); err != nil {
		panic(err)
	}
}

func main() {
	addr := "localhost:8080"

	infoLog := log.New(log.Writer(), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errLog := log.New(log.Writer(), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	db := common.InitDB("holdbillulac.db", infoLog)
	runGooseMigration(db)
	api.Initialize(db, infoLog, errLog)

	infoLog.Printf("Listening on: %s", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		errLog.Fatal(err)
	}
}
