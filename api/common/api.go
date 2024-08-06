package common

import (
	"github.com/a-h/templ"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type insertable interface {
	SetId(id int)
}

type CRUD struct {
	Insert    string
	SelectAll string
	Select    string
	Delete    string
	Update    string
}

func Get[T insertable](db *sqlx.DB, w http.ResponseWriter, selectQuery string, id string, render func(T) templ.Component) error {
	player, err := Query[T](db, selectQuery, id)
	if err != nil {
		return err
	}
	err = ToHTML[T](w, render, player)
	if err != nil {
		return err
	}
	return nil
}

func GetAll[T insertable](db *sqlx.DB, w http.ResponseWriter, query string, renderer func(T) templ.Component) error {
	items, err := Query[[]T](db, query)
	if err != nil {
		return err
	}
	err = ListToHTML[T](w, renderer, items)
	if err != nil {
		return err
	}
	return nil
}

func Create[T insertable](db *sqlx.DB, w http.ResponseWriter, query string, item T, renderer func(T) templ.Component) error {
	insertId, err := Insert(db, query, item)
	if err != nil {
		return err
	}
	item.SetId(insertId)
	err = ToHTML[T](w, renderer, item)
	if err != nil {
		return err
	}
	return nil
}
