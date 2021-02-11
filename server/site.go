package main

import (
	"encoding/json"
	// "os"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)


	// TODO: Build Struct to Handle data/* files

type Page struct {
	ContentTitle string
	// SubTitle string `json: "content.subHeader"`
	// TextBody string `json: "content.body"`
	// List []string `json: "content.listOne"`

}

type PageTwo struct {
	ContentTitle string `json:"title"`
	// SubTitle string `json: "content.subHeader"`
	// TextBody string `json: "content.body"`
	// List []string `json: "content.listOne"`

}

var templates = template.Must(template.ParseGlob("web/templates/*.html"))
var validPath = regexp.MustCompile("^/(home)/([a-zA-Z0-9]+)$")

func loadFile() {
	var p PageTwo

	file, err := ioutil.ReadFile("./data/home.json")
	if err != nil{
		//TODO: needs to change
		fmt.Println(err)
	}

	err = json.Unmarshal(file, &p)
	if err != nil {
		//TODO: needs to change
        fmt.Println("error:", err)
    }

	fmt.Println("Result:", p.ContentTitle)
	//TODO: This now sort of works. We now need to pass this around.
}

func loadPage(title string) *Page {
	// TODO: Make this actually useful. Starting block for loading from data/* files
	// TODO: Figure out how to start incorporating nested templates
	return &Page{}
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
	loadFile()
    http.HandleFunc("/home/", makeHandler(viewHandler))
    log.Fatal(http.ListenAndServe(":8080", nil))
}