package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /products", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("GET ALL PRODUCTS"))
	})

	mux.HandleFunc("GET /products/{id}", func(w http.ResponseWriter, r *http.Request) {
		productID := r.PathValue("id")

		w.Write([]byte(fmt.Sprintf("GET PRODUCT WITH ID: %s", productID)))
	})

	mux.HandleFunc("POST /products", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("CREATE A PRODUCT"))
	})

	mux.HandleFunc("PATCH /products", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("UPDATE A PRODUCT"))
	})

	log.Fatal(http.ListenAndServe(":3000", mux))

}
