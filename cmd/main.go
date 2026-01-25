package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	handlerCategory "github.com/DioSaputra28/belejar-go-dasar/internal/category/handler"
	handlerProduct "github.com/DioSaputra28/belejar-go-dasar/internal/produk/handler"

	_ "github.com/DioSaputra28/belejar-go-dasar/docs" // This line is needed for swagger
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title Product & Category API
// @version 1.0
// @description This is a simple REST API for managing products and categories
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email dio@example.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

func main() {

	// GET localhost:8080/api/produk/{id}
	// PUT localhost:8080/api/produk/{id}
	// DELETE localhost:8080/api/produk/{id}
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handlerProduct.GetProdukByID(w, r)
		} else if r.Method == "PUT" {
			handlerProduct.UpdateProduk(w, r)
		} else if r.Method == "DELETE" {
			handlerProduct.DeleteProduk(w, r)
		}

	})

	// GET localhost:8080/api/produk
	// POST localhost:8080/api/produk
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handlerProduct.GetProduk(w, r)
		} else if r.Method == "POST" {
			handlerProduct.CreateProduk(w, r)
		}

	})

	// GET localhost:8080/api/category/{id}
	// PUT localhost:8080/api/category/{id}
	// DELETE localhost:8080/api/category/{id}
	http.HandleFunc("/api/category/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handlerCategory.GetCategoryByID(w, r)
		} else if r.Method == "PUT" {
			handlerCategory.UpdateCategory(w, r)
		} else if r.Method == "DELETE" {
			handlerCategory.DeleteCategory(w, r)
		}
	})

	// GET localhost:8080/api/category
	// POST localhost:8080/api/category
	http.HandleFunc("/api/category", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handlerCategory.GetCategory(w, r)
		} else if r.Method == "POST" {
			handlerCategory.CreateCategory(w, r)
		}
	})

	// Swagger UI route
	// Access at: http://localhost:8080/swagger/index.html
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	// localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})
	fmt.Println("Server running di localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("gagal running server")
		fmt.Println(err)
	}
}
