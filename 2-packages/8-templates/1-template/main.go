package main

import (
	"html/template"
	"os"
)

type Course struct {
	Name     string
	Workload int
}

func main() {
	course := Course{
		Name:     "GOLang",
		Workload: 10,
	}
	tmp := template.New("CourseTemplate")
	tmp, _ = tmp.Parse("Course: {{.Name}} - Workload: {{.Workload}}")
	err := tmp.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}
