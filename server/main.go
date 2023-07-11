package main

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/health-check", healthCheck)
	router.Get("/index", view)
	router.Get("/output.css", styles)
	router.Post("/add-todo", addTodo)
	router.Post("/remove-todo/index/{index}", removeTodo)

	log.Default().Println("Listening on port 7007.")
	http.ListenAndServe(":7007", router)
}

type Todo struct {
	Name string
}

type PageData struct {
	Items []Todo
}

var todoList []Todo = []Todo{}
var tmpl = template.Must(template.ParseGlob("server/view/*.html"))

func styles(w http.ResponseWriter, r *http.Request) {
	styles := template.Must(template.ParseFiles("server/view/styles/output.css"))
	styles.Execute(w, nil)
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	todoList = append(todoList, Todo{Name: name})

	data := PageData{Items: todoList}

	tmpl.ExecuteTemplate(w, "todo-items", data)
}

func removeTodo(w http.ResponseWriter, r *http.Request) {
	index, err := strconv.ParseInt(chi.URLParam(r, "index"), 10, 64)
	if err != nil {
		// TODO return 400
		return
	}

	todoList = append(todoList[:index], todoList[index+1:]...)

	data := PageData{Items: todoList}
	tmpl.ExecuteTemplate(w, "todo-items", data)
}

func view(w http.ResponseWriter, r *http.Request) {
	data := PageData{Items: todoList}

	tmpl.ExecuteTemplate(w, "index.html", data)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
