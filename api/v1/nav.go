package v1

import (
	"errors"
	"github.com/gorilla/mux"
	"holdbillulac/api/common"
	"net/http"
)

type Nav struct {
	common.Base
	Name string       `db:"name"`
	Data common.JSONB `db:"nav_data"`
}

var navQueries = common.CRUD{
	SelectAll: "SELECT * FROM nav",
	Select:    "SELECT * FROM nav WHERE id = ?",
	Insert:    "<NOT IMPLEMENTED>",
	Delete:    "<NOT IMPLEMENTED>",
	Update:    "<NOT IMPLEMENTED>",
}

func getNav(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		common.HandleError(w, errors.New("must supply ID"), http.StatusBadRequest)
		return
	}
	err := common.Get[Nav](db, w, navQueries.Select, id, navDiv)
	if err != nil {
		common.HandleError(w, err, http.StatusInternalServerError)
		return
	}
}

func getNavs(w http.ResponseWriter, _ *http.Request) {
	err := common.GetAll[Nav](db, w, navQueries.SelectAll, navDiv)
	if err != nil {
		common.HandleError(w, err, http.StatusInternalServerError)
		return
	}
}
