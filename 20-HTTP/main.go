package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type user struct {
	Name  string // Name is a field of the user struct and must be uppercase to be exported in the HTML template
	Age   int8
	Email string
}

var templates *template.Template

func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Just a simple root page"))
}

func home(w http.ResponseWriter, r *http.Request) {
	u := user{
		Name:  "John Wick", // Name is a field of the user struct and must be uppercase to be exported in the HTML template
		Age:   50,
		Email: "jhon@wick.com",
	}
	templates.ExecuteTemplate(w, "home.html", u) // ExecuteTemplate() executes the template associated with the given name in the set.
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Load Users page!"))
}

func main() {

	templates = template.Must(template.ParseGlob("www/*.html")) // ParseGlob() returns a template set containing all the templates matching the specified pattern.

	http.HandleFunc("/", root)

	http.HandleFunc("/home", home)

	http.HandleFunc("/users", users)

	fmt.Println("Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
