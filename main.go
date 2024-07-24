package main

import (
	rice "github.com/GeertJohan/go.rice"
	"html/template"
	"log"
	"net/http"
	"time"
)

var (
	templates = map[string]TemplateConfig{}
)

type TemplateConfig struct {
	Name       string
	Path       string
	Content    string
	HandleFunc func(w http.ResponseWriter, r *http.Request)
}

func index(w http.ResponseWriter, _ *http.Request) {
	t := templates["index"]
	msg, err := template.New(t.Name).Parse(t.Content)
	if err != nil {
		log.Fatal(err)
	}
	err = msg.Execute(w, map[string]string{"Message": time.Now().String()})
	if err != nil {
		log.Fatal(err)
	}
}

func initTemplates() {
	box := rice.MustFindBox("templates")
	templates = map[string]TemplateConfig{
		"index": {
			Path:       "GET /",
			HandleFunc: index,
			Content:    box.MustString("index.html"),
		},
	}
	for k, v := range templates {
		http.HandleFunc(v.Path, v.HandleFunc)
		log.Printf("Loaded template: %v\n", k)
	}
}

func main() {
	initTemplates()
	log.Printf("Listening on: %s", "localhost:8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
