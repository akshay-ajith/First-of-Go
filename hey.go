package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Hello yu")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		//io.WriteString(w, "Hello you\n")
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "The Godfather", Director: "Francis Ford Coppola"},
				{Title: "Blade Runner", Director: "Ridley Scott"},
			},
		}

		tmpl.Execute(w, films)

	}
	h2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)

		title := r.PostFormValue("title")
		director := r.PostFormValue("Director")
		fmt.Println(title)
		fmt.Println(director)

		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)

	log.Fatal(http.ListenAndServe(":5000", nil))

}
