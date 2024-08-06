package common

import (
	"context"
	"fmt"
	"github.com/a-h/templ"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"strconv"
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
	err = StructToHTML[T](w, render, player)
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
	err = ArrayToHTML[T](w, renderer, items)
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
	err = StructToHTML[T](w, renderer, item)
	if err != nil {
		return err
	}
	return nil
}

func ArrayToHTML[T any](w http.ResponseWriter, renderer func(T) templ.Component, items []T) error {
	for i := range items {
		err := renderer(items[i]).Render(context.Background(), w)
		if err != nil {
			return err
		}
	}
	return nil
}

func StructToHTML[T any](w http.ResponseWriter, renderer func(T) templ.Component, item T) error {
	err := renderer(item).Render(context.Background(), w)
	if err != nil {
		return err
	}
	return nil
}

func HandleError(l *log.Logger, w http.ResponseWriter, err error, errCode int) {
	l.Println(err.Error())
	http.Error(w, err.Error(), errCode)
}

func FieldToInt(raw interface{}) (int, error) {
	if mmrStr, ok := raw.(string); ok {
		if mmr, err := strconv.Atoi(mmrStr); err == nil {
			return mmr, nil
		} else {
			return 0, fmt.Errorf("field is not a valid integer")
		}
	} else {
		return 0, fmt.Errorf("field is not a number or string")
	}
}
