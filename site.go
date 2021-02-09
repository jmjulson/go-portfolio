package main

import (
	"html/template"
	"log"
	"net/http"
	"regexp"
)

type Page struct {
	// TODO: Build Struct to Handle data/* files
	PageTitle string
}

var templates = template.Must(template.ParseGlob("web/templates/*.html"))
var validPath = regexp.MustCompile("^/(home)/([a-zA-Z0-9]+)$")

func loadPage(title string) *Page {
	return &Page{PageTitle: title}
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p := loadPage("Page Name")
	// TODO: get information from data/*
	// TODO: Add proper error handling
    renderTemplate(w, "home.html", p)
}

func makeHandler(fn func(ww http.ResponseWriter, rr *http.Request, text string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := "text"
		fn(w, r, m)
	}
}

func main() {
    //tmpl := template.Must(template.ParseFiles("web/templates/home.html"))
    http.HandleFunc("/home/", makeHandler(viewHandler))
    log.Fatal(http.ListenAndServe(":8080", nil))
}