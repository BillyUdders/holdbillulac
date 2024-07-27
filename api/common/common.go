package common

import (
	"log"
	"net/http"
)

type Base struct {
	ID int64 `db:"id"`
}

func HandleError(l *log.Logger, w http.ResponseWriter, err error, errCode int) {
	l.Println(err.Error())
	http.Error(w, err.Error(), errCode)
}
