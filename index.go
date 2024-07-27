package main

import (
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, _ *http.Request) {
	content := box.MustString("index.html")
	tmpl := template.Must(template.New("index").Parse(content))
	err := tmpl.Execute(w, indexCtx)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
}
