package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
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
		title := r.PostFormValue("title")
		director := r.PostFormValue("Director")
		fmt.Println(title)
		fmt.Println(director)
		htmlstr := fmt.Sprintf("<li class = 'list-group-item bg-primary text-white'> %s - %s</li>", title, director)
		tmpl, _ := template.New("t").Parse(htmlstr)
		tmpl.Execute(w, nil)
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)

	log.Fatal(http.ListenAndServe(":5000", nil))

}
