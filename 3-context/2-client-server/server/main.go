package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(writer http.ResponseWriter, request *http.Request) {
	cxt := request.Context()
	log.Println("Request Initialized")
	defer log.Println("Request completed")
	select {
	case <-time.After(5 * time.Second):
		log.Println("Request proccessed")
		writer.Write([]byte("Request proccessed successfully"))
	case <-cxt.Done():
		writer.Write([]byte("Request canceled by client"))
	}
}
