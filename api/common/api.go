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

type mapper[T any] func(T) templ.Component

type CRUD struct {
	Insert    string
	SelectAll string
	Select    string
	Delete    string
	Update    string
}

func Get[T insertable](db *sqlx.DB, w http.ResponseWriter, selectQuery string, id string, tMap mapper[T]) error {
	player, err := Query[T](db, selectQuery, id)
	if err != nil {
		return err
	}
	err = StructHTMLResponse[T](w, player, tMap)
	if err != nil {
		return err
	}
	return nil
}

func GetAll[T insertable](db *sqlx.DB, w http.ResponseWriter, query string, tMap mapper[T]) error {
	items, err := Query[[]T](db, query)
	if err != nil {
		return err
	}
	err = ArrayHTMLResponse[T](w, items, tMap)
	if err != nil {
		return err
	}
	return nil
}

func Create[T insertable](db *sqlx.DB, w http.ResponseWriter, query string, item T, tMap mapper[T]) error {
	insertId, err := Insert(db, query, item)
	if err != nil {
		return err
	}
	item.SetId(insertId)
	err = StructHTMLResponse[T](w, item, tMap)
	if err != nil {
		return err
	}
	return nil
}

func ArrayHTMLResponse[T any](w http.ResponseWriter, items []T, tMap mapper[T]) error {
	for i := range items {
		err := tMap(items[i]).Render(context.Background(), w)
		if err != nil {
			return err
		}
	}
	return nil
}

func StructHTMLResponse[T any](w http.ResponseWriter, item T, tMap mapper[T]) error {
	err := tMap(item).Render(context.Background(), w)
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
