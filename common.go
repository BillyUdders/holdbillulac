package main

import "net/http"

type Base struct {
	ID int64 `db:"id"`
}

func handleError(w http.ResponseWriter, err error, errCode int) {
	errLog.Println(err.Error())
	http.Error(w, err.Error(), errCode)
}
