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
	// TODO: Make this actually useful. Starting block for loading from data/* files
	// TODO: Figure out how to start incorporating nested templates
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
	// TODO: get actual information from loadPage
	// TODO: Add proper error handling
    renderTemplate(w, "home.html", p)
}

func makeHandler(fn func(ww http.ResponseWriter, rr *http.Request, text string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Figure out what to do about this...
		m := "text"
		fn(w, r, m)
	}
}

func main() {
	// TODO: make the first param here dynamic. Make it file driven.
    http.HandleFunc("/home/", makeHandler(viewHandler))
    log.Fatal(http.ListenAndServe(":8080", nil))
}