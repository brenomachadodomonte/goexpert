package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	//mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	writer.Write([]byte("Hello World!"))
	//})
	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blog{title: "Blog Page!"})

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
	println("Listening...")
}

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello World!"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(b.title))
}
