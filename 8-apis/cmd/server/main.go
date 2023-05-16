package main

import (
	"github.com/brenomachadodomonte/goexpert/apis/configs"
	"github.com/brenomachadodomonte/goexpert/apis/internal/entity"
	"github.com/brenomachadodomonte/goexpert/apis/internal/infra/database"
	"github.com/brenomachadodomonte/goexpert/apis/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(entity.Product{}, entity.User{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8000", nil)
}
