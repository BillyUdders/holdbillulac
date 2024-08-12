package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Base struct {
	ID int `db:"id"`
}

func (b *Base) SetId(id int) {
	b.ID = id
	fmt.Println()
}

type JSONB map[string]interface{}

func (p *JSONB) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *JSONB) Scan(src interface{}) error {
	source, ok := src.(string)
	if !ok {
		return errors.New("type assertion .(string) failed")
	}
	if err := json.Unmarshal([]byte(source), p); err != nil {
		return err
	}
	return nil
}

func InitDB(dbName string) *sqlx.DB {
	return sqlx.MustConnect("sqlite3", dbName)
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
