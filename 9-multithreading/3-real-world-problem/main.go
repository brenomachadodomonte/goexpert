package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var number uint64 = 0

// go run -race main.go to check if is there a race condition
func main() {
	//m := sync.Mutex{}
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		//m.Lock()
		atomic.AddUint64(&number, 1)
		//m.Unlock()
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
