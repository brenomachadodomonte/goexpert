package main

import (
	"log"
	"net/http"
)

func recoverHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Println("panic recovered", r)
				http.Error(writer, "Internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(writer, request)
	})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World"))
	})

	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		panic("panic")
	})

	log.Fatal(http.ListenAndServe(":3000", recoverHandler(mux)))
}
