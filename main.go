package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Note struct {
	Id       uint16 `gorm:"primaryKey"` // Добавить тег для первичного ключа
	Title    string
	FullText string
}

var db *gorm.DB
var notes []Note
var showNote Note

func initDB() {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to connect to database: %v", err))
	}

	err = db.AutoMigrate(&Note{})
	if err != nil {
		panic(fmt.Errorf("failed to migrate database: %v", err))
	}

	fmt.Println("Successfully connected to database!")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tmp/footer.html", "tmp/header.html", "tmp/home_page.html")
	if err != nil {
		http.Error(w, "Ошибка загрузки шаблона", http.StatusInternalServerError)
		return
	}

	// Получение всех заметок из базы данных
	if err := db.Find(&notes).Error; err != nil {
		http.Error(w, "Ошибка получения заметок", http.StatusInternalServerError)
		return
	}

	t.ExecuteTemplate(w, "home_page", notes)
}

func create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tmp/footer.html", "tmp/header.html", "tmp/create.html")
	if err != nil {
		http.Error(w, "Ошибка загрузки шаблона", http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w, "create", nil)
}

func saveNote(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	fullText := r.FormValue("full_text")

	// Создание новой заметки и сохранение в базе данных
	note := Note{Title: title, FullText: fullText}
	if err := db.Create(&note).Error; err != nil {
		http.Error(w, "Ошибка сохранения заметки", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteNote(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	noteId := r.FormValue("id")

	// Удаление заметки из базы данных
	if err := db.Delete(&Note{}, noteId).Error; err != nil {
		http.Error(w, "Ошибка удаления заметки", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func showNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	t, err := template.ParseFiles("tmp/show.html", "tmp/header.html", "tmp/footer.html")
	if err != nil {
		http.Error(w, "Ошибка загрузки шаблона", http.StatusInternalServerError)
		return
	}

	var note Note

	// Получение заметки по ID
	if err := db.First(&note, vars["id"]).Error; err != nil {
		http.Error(w, "Ошибка получения заметки", http.StatusInternalServerError)
		return
	}

	showNote = note
	t.ExecuteTemplate(w, "show", showNote)
}

func title(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tmp/title.html", "tmp/header.html", "tmp/footer.html")
	if err != nil {
		http.Error(w, "Ошибка загрузки шаблона", http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w, "title", nil)
}

func handleFunc() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", homePage).Methods("GET")
	rtr.HandleFunc("/create", create).Methods("GET")
	rtr.HandleFunc("/save_note", saveNote).Methods("POST")
	rtr.HandleFunc("/note/{id:[0-9]+}", showNoteHandler).Methods("GET")
	rtr.HandleFunc("/delete_note", deleteNote).Methods("POST")
	rtr.HandleFunc("/title", title).Methods("GET")

	http.Handle("/", rtr)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Println("Starting server...")
	initDB() // Инициализация базы данных
	handleFunc()
}
