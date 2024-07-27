package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB(dbName string) *sqlx.DB {
	val, err := sqlx.Connect("sqlite3", dbName)
	if err != nil {
		log.Fatalln(err)
	}
	return val
}

func Insert(db *sqlx.DB, query string, insertable any) (int64, error) {
	stmt, err := db.PrepareNamed(query)
	if err != nil {
		return 0, err
	}
	id, err := stmt.Exec(insertable)
	if err != nil {
		return 0, err
	}
	insertId, err := id.LastInsertId()
	if err != nil {
		return 0, err
	}
	return insertId, err
}

func Query[T any](db *sqlx.DB, query string, params ...interface{}) (T, error) {
	var all T
	err := db.Select(&all, query, params...)
	return all, err
}
