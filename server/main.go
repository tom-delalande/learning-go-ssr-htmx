package main

import (
	"log"
	"net/http"
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

	log.Default().Println("Listening on port 7007.")
	http.ListenAndServe(":7007", router)
}

func styles(w http.ResponseWriter, r *http.Request) {
	type PageData struct {
	}
	tmpl := template.Must(template.ParseFiles("view/styles/output.css"))

	tmpl.Execute(w, PageData{})
}

func view(w http.ResponseWriter, r *http.Request) {
	type PageData struct {
		Name string
	}
	tmpl := template.Must(template.ParseFiles("view/index.html"))

	tmpl.Execute(w, PageData{
		Name: "Lalande",
	})
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
