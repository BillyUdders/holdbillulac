package common

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func InitDB(dbName string, createTableStmt string, logger *log.Logger) *sqlx.DB {
	db := sqlx.MustConnect("sqlite3", dbName)
	_ = db.MustExec(createTableStmt)
	logger.Println("SQLLite3 Database: initialized")
	return db
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
