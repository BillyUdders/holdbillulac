package v1

//
//import (
//	"errors"
//	"github.com/gorilla/mux"
//	"holdbillulac/api/common"
//	"net/http"
//)
//
////var navConfig = CRUDConfig{
////	Insert:    "INSERT INTO players (name, age, MMR) VALUES (:name, :age, :MMR)",
////	SelectAll: "SELECT * FROM players",
////	Select:    "SELECT * FROM players WHERE id = ?",
////	Delete:    "DELETE FROM players WHERE id = ?",
////	Update:    "UPDATE QUERY",
////	Template: template.Must(template.New("navbar-item").Parse(`
////        <div class="dropdown">
////            <button class="dropbtn">Dropdown 2</button>
////            <div class="dropdown-content">
////                <a href="#">Link 4</a>
////                <a href="#">Link 5</a>
////                <a href="#">Link 6</a>
////            </div>
////        </div>
////	`)),
////}
//
//func getNav(w http.ResponseWriter, r *http.Request) {
//	id := mux.Vars(r)["id"]
//	if id == "" {
//		common.HandleError(errLog, w, errors.New("must supply ID"), http.StatusBadRequest)
//		return
//	}
//	player, err := common.Query[Nav](db, selectByID, id)
//	if err != nil {
//		common.HandleError(errLog, w, err, http.StatusInternalServerError)
//		return
//	}
//	err = trTemplate.Execute(w, player)
//	if err != nil {
//		common.HandleError(errLog, w, err, http.StatusInternalServerError)
//		return
//	}
//	infoLog.Printf("Get all: %v", player)
//}
//
//func getNavs(w http.ResponseWriter, _ *http.Request) {
//	players, err := common.Query[[]Nav](db, selectAll)
//	if err != nil {
//		common.HandleError(errLog, w, err, http.StatusInternalServerError)
//		return
//	}
//	for i := range players {
//		player := players[i]
//		err = trTemplate.Execute(w, player)
//		if err != nil {
//			common.HandleError(errLog, w, err, http.StatusInternalServerError)
//			return
//		}
//	}
//	infoLog.Printf("Navs returned: %v", len(players))
//}
//
//func createNav(w http.ResponseWriter, r *http.Request) {
//	player, err := new(Nav).fromBody(r.Body)
//	if err != nil {
//		common.HandleError(errLog, w, err, http.StatusBadRequest)
//		return
//	}
//	insertId, err := common.Insert(db, insert, player)
//	if err != nil {
//		common.HandleError(errLog, w, err, http.StatusInternalServerError)
//		return
//	}
//	player.ID = insertId
//	err = trTemplate.Execute(w, player)
//	if err != nil {
//		common.HandleError(errLog, w, err, http.StatusInternalServerError)
//		return
//	}
//	infoLog.Printf("Created: %v", player)
//}
//
//func deleteNav(w http.ResponseWriter, r *http.Request) {
//	id := mux.Vars(r)["id"]
//	_, err := db.Exec(deleteByID, id)
//	if err != nil {
//		common.HandleError(errLog, w, err, http.StatusInternalServerError)
//		return
//	}
//	infoLog.Printf("Deleted ID: %v", id)
//}
