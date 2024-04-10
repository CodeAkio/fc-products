package main

import (
	"github.com/CodeAkio/fc-products/configs"
	"github.com/CodeAkio/fc-products/internal/entity"
	"github.com/CodeAkio/fc-products/internal/infra/database"
	"github.com/CodeAkio/fc-products/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
		return
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&entity.Product{}, &entity.User{})
	if err != nil {
		panic(err)
		return
	}
	productDB := database.NewProduct(db)

	productHandler := handlers.NewProductHandler(productDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)

	err = http.ListenAndServe(":8000", r)
	if err != nil {
		panic(err)
		return
	}
}
