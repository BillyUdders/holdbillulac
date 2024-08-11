package main

import (
	"embed"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	"holdbillulac/api/common"
	api "holdbillulac/api/v1"
	"log/slog"
	"net/http"
	"os"
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

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	db := common.InitDB("holdbillulac.db")
	runGooseMigration(db)
	router := api.Initialize(db)

	logger.Info("Listening:", "address", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		logger.Error("", err)
	}
}
