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

func defaultLoggingConfig() {
	logLevel := &slog.LevelVar{}
	logLevel.Set(slog.LevelInfo)
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))
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
	defaultLoggingConfig()
	addr := ":8080"
	dbName := "holdbillulac.db"

	db := common.InitDB(dbName)
	runGooseMigration(db)
	router := api.Initialize(db)

	slog.Info("Initialized", "address", addr, "dbName", dbName)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		slog.Error("Server error", err)
	}
}
