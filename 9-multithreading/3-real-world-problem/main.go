package main

import (
	"fmt"
	"net/http"
	"sync"
)

var number uint64 = 0

func main() {
	m := sync.Mutex{}
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		m.Lock()
		number++
		m.Unlock()
		_, err := writer.Write([]byte(fmt.Sprintf("You accessed this page %d times", number)))
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		panic(err)
	}
}
