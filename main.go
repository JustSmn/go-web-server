package main

import (
	"fmt"
	"html/template"
	"net/http"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Note struct {
	Id              uint16
	Title, FullText string
}

var notes = []Note{}
var showNote = Note{}

func home_page(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello World!")
	t, err := template.ParseFiles("tmp/footer.html", "tmp/header.html", "tmp/home_page.html")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("mysql", "mysql:mysql@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	res, err := db.Query("SELECT * FROM `notes`")
	if err != nil {
		panic(err)
	}

	notes = []Note{}
	for res.Next() {
		var note Note
		err = res.Scan(&note.Id, &note.Title, &note.FullText)
		if err != nil {
			panic(err)
		}
		notes = append(notes, note)
	}
	t.ExecuteTemplate(w, "home_page", notes)

}

func create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tmp/footer.html", "tmp/header.html", "tmp/create.html")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(w, "create", nil)
}

func save_note(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	full_text := r.FormValue("full_text")

	db, err := sql.Open("mysql", "mysql:mysql@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	insert, err := db.Query(fmt.Sprintf("INSERT INTO `notes` (`Title`, `FullText`) VALUES('%s', '%s')", title, full_text))
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func delete_note(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // Парсинг формы
	noteId := r.FormValue("id")
	//noteId := r.URL.Path[len("/note/"):]

	db, err := sql.Open("mysql", "mysql:mysql@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	delete, err := db.Exec(fmt.Sprintf("DELETE FROM `notes` WHERE `id` = '%s'", noteId))
	if err != nil {
		panic(err)
	}
	fmt.Println(delete)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func show_note(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	t, err := template.ParseFiles("tmp/show.html", "tmp/header.html", "tmp/footer.html")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("mysql", "mysql:mysql@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	res, err := db.Query(fmt.Sprintf("SELECT * FROM `notes` WHERE `id` = '%s' ", vars["id"]))
	if err != nil {
		panic(err)
	}

	showNote = Note{}

	for res.Next() {
		var post Note
		err = res.Scan(&post.Id, &post.Title, &post.FullText)
		if err != nil {
			panic(err)
		}

		showNote = post
	}

	t.ExecuteTemplate(w, "show", showNote)
}

func title(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tmp/title.html", "tmp/header.html", "tmp/footer.html")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "title", nil)

}

func handleFunc() {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/", home_page).Methods("GET")
	rtr.HandleFunc("/create", create).Methods("GET")
	rtr.HandleFunc("/save_note", save_note).Methods("POST")
	rtr.HandleFunc("/note/{id:[0-9]+}", show_note).Methods("GET")
	rtr.HandleFunc("/delete_note", delete_note).Methods("POST")
	rtr.HandleFunc("/title", title).Methods("GET")

	http.Handle("/", rtr)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	http.ListenAndServe(":8080", nil)
}

func main() {
	handleFunc()
}
