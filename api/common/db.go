package common

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Base struct {
	ID int `db:"id"`
}

func (b *Base) SetId(id int) {
	b.ID = id
	fmt.Println()
}

func InitDB(dbName string, logger *log.Logger) *sqlx.DB {
	db := sqlx.MustConnect("sqlite3", dbName)
	logger.Printf("Initialized: %s", dbName)
	return db
}

func Insert(db *sqlx.DB, query string, insertable any) (int, error) {
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
	return int(insertId), err
}

func Query[T any](db *sqlx.DB, query string, params ...interface{}) (T, error) {
	var all T
	err := db.Select(&all, query, params...)
	return all, err
}
