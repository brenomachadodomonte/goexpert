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
		t := template.Must(template.New("index.html").ParseFiles("index.html"))
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
