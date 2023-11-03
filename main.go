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
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film {
			"Films" : {
				{Title: "The Matrix", Director: "Wachowski"},
				{Title: "The Matrix Reloaded", Director: "Wachowski"},
				{Title: "The Matrix Revolutions", Director: "Wachowski"},
			},
		}
		tmpl.Execute(w, films)
	}

	h2 := func (w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		htmlStr := fmt.Sprintf("<li class='w-full px-4 py-2 border-b border-gray-200 rounded-t-lg dark:border-gray-600'> %s - %s</li>", title, director)
		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, nil)
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/film", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
