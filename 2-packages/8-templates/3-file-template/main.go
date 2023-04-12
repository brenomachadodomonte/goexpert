package main

import (
	"html/template"
	"os"
)

type Course struct {
	Name     string
	Workload int
}

type Courses []Course

func main() {
	t := template.Must(template.New("index.html").ParseFiles("index.html"))
	err := t.Execute(os.Stdout, Courses{
		{"Go", 10},
		{"Java", 8},
		{"Python", 5},
		{"React", 4},
	})
	if err != nil {
		panic(err)
	}
}
