/*
There is a business logic realization

	Processing HTTP requests
	Working with models
	Forming a response
	Handles errors
*/
package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"main.go/storage"
)

type NoteController struct {
	db *gorm.DB
}

func NewNoteController(db *gorm.DB) *NoteController {
	return &NoteController{db: db}
}

func (h *NoteController) HomePage(w http.ResponseWriter, r *http.Request) {
	notes, err := storage.GetAllNotes(h.db)
	if err != nil {
		http.Error(w, "Ошибка получения заметок", http.StatusInternalServerError)
		return
	}
	renderTemplate(w, "home_page", notes)
}

func (c *NoteController) CreatePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "create", nil)
}

func (c *NoteController) SaveNote(w http.ResponseWriter, r *http.Request) {
	note := &storage.Note{
		Title:    r.FormValue("title"),
		FullText: r.FormValue("full_text"),
	}

	if err := storage.CreateNote(note); err != nil {
		http.Error(w, "Ошибка сохранения заметки", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (c *NoteController) DeleteNote(w http.ResponseWriter, r *http.Request) {
	if err := storage.DeleteNote(r.FormValue("id")); err != nil {
		http.Error(w, "Ошибка удаления заметки", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (c *NoteController) ShowNote(w http.ResponseWriter, r *http.Request) {
	note, err := storage.GetNoteByID(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Ошибка получения заметки", http.StatusInternalServerError)
		return
	}
	renderTemplate(w, "show", note)
}

func (c *NoteController) TitlePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "title", nil)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles(
		filepath.Join("tmp", "header.html"),
		filepath.Join("tmp", "footer.html"),
		filepath.Join("tmp", tmpl+".html"),
	)
	if err != nil {
		http.Error(w, "Ошибка загрузки шаблона: "+err.Error(), http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w, tmpl, data)
}
