package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"
)

var (
	editHTMLFile = "edit.html"
	viewHTMLFile = "view.html"
	templates    = template.Must(template.ParseFiles(editHTMLFile, viewHTMLFile))
)

type Page struct {
	Title string
	Body  string
}

func viewHandler() {
	content := &Page{
		Title: "View Page",
		Body:  "This is view page",
	}
	renderTemplate(viewHTMLFile, content)
}

func editHandler() {
	content := &Page{
		Title: "Edit Page",
		Body:  "This is edit page",
	}
	renderTemplate(editHTMLFile, content)
}

func renderTemplate(htmlFile string, content *Page) {
	if err := templates.ExecuteTemplate(os.Stdout, htmlFile, content); err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

func main() {
	viewHandler()

	fmt.Println()
	fmt.Println()
	fmt.Println(strings.Repeat("-", 100))
	fmt.Println()

	editHandler()
}
