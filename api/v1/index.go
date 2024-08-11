package v1

import (
	"context"
	"holdbillulac/api/common"
	"net/http"
)

type IndexPage struct {
	Title       string
	Description string
	Tagline     string
}

var idx = IndexPage{
	Title:       "Billy Stats",
	Description: "BILLIGULAC",
	Tagline:     "Unashamed clone of Aligulac for fun and learning",
}

func index(w http.ResponseWriter, _ *http.Request) {
	if err := indexPage(idx).Render(context.Background(), w); err != nil {
		common.HandleError(w, err, http.StatusInternalServerError)
		return
	}
}
