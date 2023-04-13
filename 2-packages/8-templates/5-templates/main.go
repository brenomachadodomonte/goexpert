package main

import (
	"html/template"
	"net/http"
)

type Course struct {
	Name     string
	Workload int
}

type Courses []Course

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		templates := []string{
			"header.html",
			"content.html",
			"footer.html",
		}
		t := template.Must(template.New("content.html").ParseFiles(templates...))
		err := t.Execute(writer, Courses{
			{"Go", 10},
			{"Java", 8},
			{"Python", 5},
			{"React", 4},
		})
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8080", nil)
}
