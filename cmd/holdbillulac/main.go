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

func configSlog() {
	logLevel := &slog.LevelVar{}
	logLevel.Set(slog.LevelDebug)
	opts := &slog.HandlerOptions{Level: logLevel}
	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))
	slog.SetDefault(logger)
}

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
	configSlog()

	addr := ":8080"
	dbName := "holdbillulac.db"

	db := common.InitDB(dbName)
	runGooseMigration(db)
	router := api.Initialize(db)
	slog.Info("Initialized", "address", addr, "dbName", dbName)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}
