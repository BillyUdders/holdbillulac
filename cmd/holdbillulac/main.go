package main

import (
	api "holdbillulac/api/v1"
	"log"
	"net/http"
)

func main() {
	addr := "localhost:8080"

	infoLog := log.New(log.Writer(), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errLog := log.New(log.Writer(), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	api.Initialize("holdbillulac.db", infoLog, errLog)

	infoLog.Printf("Listening on: %s", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		errLog.Fatal(err)
	}
}
