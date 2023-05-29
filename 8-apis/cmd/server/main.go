package main

import (
	"github.com/brenomachadodomonte/goexpert/apis/configs"
	"github.com/brenomachadodomonte/goexpert/apis/internal/entity"
	"github.com/brenomachadodomonte/goexpert/apis/internal/infra/database"
	"github.com/brenomachadodomonte/goexpert/apis/internal/infra/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	config, err := configs.LoadConfig(".")
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

	userDB := database.UserDB{DB: db}
	userHandler := handlers.NewUserHandler(userDB, config.TokenAuth, config.JwtExpiresIn)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products", productHandler.GetProducts)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)

	r.Post("/users", userHandler.Create)
	r.Post("/users/generate_token", userHandler.GetJWT)

	http.ListenAndServe(":8000", r)
}
