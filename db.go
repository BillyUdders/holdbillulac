package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func initDB(dbName string) *sqlx.DB {
	val, err := sqlx.Connect("sqlite3", dbName)
	if err != nil {
		log.Fatalln(err)
	}
	return val
}

func insert(db *sqlx.DB, query string, params ...interface{}) (int64, error) {
	exec, err := db.Exec(query, params...)
	if err != nil {
		return 0, err
	}
	id, err := exec.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, err
}

func Query[T any](db *sqlx.DB, query string) (T, error) {
	var all T
	db.
	err := db.Select(&all, query)
	return all, err
}
