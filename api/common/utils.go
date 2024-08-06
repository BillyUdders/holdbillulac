package common

import (
	"context"
	"fmt"
	"github.com/a-h/templ"
	"log"
	"net/http"
	"strconv"
)

func ListToHTML[T any](w http.ResponseWriter, renderer func(T) templ.Component, items []T) error {
	for i := range items {
		err := renderer(items[i]).Render(context.Background(), w)
		if err != nil {
			return err
		}
	}
	return nil
}

func ToHTML[T any](w http.ResponseWriter, renderer func(T) templ.Component, item T) error {
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
