package v1

import (
	"holdbillulac/api/common"
	"html/template"
	"net/http"
)

var indexCtx = map[string]string{
	"Title":       "Big Test Page",
	"Description": "Holden x Bill: Aligulac",
}

func Index(w http.ResponseWriter, _ *http.Request) {
	content := box.MustString("index.html")
	tmpl := template.Must(template.New("index").Parse(content))
	err := tmpl.Execute(w, indexCtx)
	if err != nil {
		common.HandleError(errLog, w, err, http.StatusInternalServerError)
		return
	}
}
