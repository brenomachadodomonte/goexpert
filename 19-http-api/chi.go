// package main
//
// import (
//
//	"fmt"
//	"github.com/go-chi/chi"
//	"log"
//	"net/http"
//
// )
//
//	func main() {
//		router := chi.NewRouter()
//
//		router.Get("/products", func(writer http.ResponseWriter, request *http.Request) {
//			writer.WriteHeader(http.StatusOK)
//			writer.Write([]byte("GET ALL PRODUCTS"))
//		})
//
//		router.Get("/products/{id}", func(writer http.ResponseWriter, request *http.Request) {
//			productID := chi.URLParam(request, "id")
//
//			writer.WriteHeader(http.StatusOK)
//			writer.Write([]byte(fmt.Sprintf("GET PRODUCT WITH ID: %s", productID)))
//		})
//
//		router.Post("/products", func(writer http.ResponseWriter, request *http.Request) {
//			writer.WriteHeader(http.StatusOK)
//			writer.Write([]byte("CREATE A PRODUCT"))
//		})
//
//		log.Fatal(http.ListenAndServe(":3000", router))
//	}
package main

func main() {

}
