/*

	There is a routing logic:
		Configuring application URLs paths
		Binding paths to handlers

*/

package routes

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"main.go/handlers"
	"net/http"
)

func Setup(db *gorm.DB) *mux.Router {
	r := mux.NewRouter()
	nc := handlers.NewNoteController(db)

	r.HandleFunc("/", nc.HomePage).Methods("GET")
	r.HandleFunc("/create", nc.CreatePage).Methods("GET")
	r.HandleFunc("/save_note", nc.SaveNote).Methods("POST")
	r.HandleFunc("/note/{id:[0-9]+}", nc.ShowNote).Methods("GET")
	r.HandleFunc("/delete_note", nc.DeleteNote).Methods("POST")
	r.HandleFunc("/title", nc.TitlePage).Methods("GET")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	return r
}
