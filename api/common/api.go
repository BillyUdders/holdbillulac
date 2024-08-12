package common

import (
	"bytes"
	"context"
	"fmt"
	"github.com/a-h/templ"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"net/http"
	"reflect"
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
	item, err := Query[T](db, selectQuery, id)
	if err != nil {
		return err
	}
	if err = structHTMLResponse[T](w, item, tMap); err != nil {
		return err
	}
	slog.Info("Get", "type", reflect.TypeOf(item), "id", id)
	return nil
}

func GetAll[T insertable](db *sqlx.DB, w http.ResponseWriter, query string, tMap mapper[T]) error {
	items, err := Query[[]T](db, query)
	if err != nil {
		return err
	}
	if err = arrayHTMLResponse[T](w, items, tMap); err != nil {
		return err
	}
	slog.Info("GetAll", "type", reflect.TypeOf(items[0]), "count", len(items))
	return nil
}

func Create[T insertable](db *sqlx.DB, w http.ResponseWriter, query string, item T, tMap mapper[T]) error {
	insertId, err := Insert(db, query, item)
	if err != nil {
		return err
	}
	item.SetId(insertId)
	if err = structHTMLResponse[T](w, item, tMap); err != nil {
		return err
	}
	slog.Info("Create", "type", reflect.TypeOf(item), "id", insertId)
	return nil
}

func HandleError(w http.ResponseWriter, err error, errCode int) {
	slog.Error(err.Error())
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

func arrayHTMLResponse[T any](w http.ResponseWriter, items []T, tMap mapper[T]) error {
	for _, val := range items {
		if err := structHTMLResponse[T](w, val, tMap); err != nil {
			return err
		}
	}
	return nil
}

func structHTMLResponse[T any](w http.ResponseWriter, item T, tMap mapper[T]) error {
	ctx := context.Background()
	component := tMap(item)
	if err := structLog(ctx, component, item); err != nil {
		return err
	}
	if err := component.Render(ctx, w); err != nil {
		return err
	}
	return nil
}

func structLog(ctx context.Context, component templ.Component, item interface{}) error {
	if slog.Default().Handler().Enabled(ctx, slog.LevelDebug) {
		var buff bytes.Buffer
		err := component.Render(ctx, &buff)
		if err != nil {
			return err
		}
		slog.Debug(
			"Render",
			"type", fmt.Sprintf("%T", item),
			"item", item,
			"rendered_html", buff.String(),
		)
		return nil
	}
	return nil
}
