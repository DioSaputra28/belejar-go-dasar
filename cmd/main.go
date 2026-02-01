package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/DioSaputra28/belejar-go-dasar/database"
	handlerCategory "github.com/DioSaputra28/belejar-go-dasar/internal/category/handler"
	categoryRepository "github.com/DioSaputra28/belejar-go-dasar/internal/category/repository"
	categoryService "github.com/DioSaputra28/belejar-go-dasar/internal/category/service"
	handlerProduct "github.com/DioSaputra28/belejar-go-dasar/internal/produk/handler"
	productRepository "github.com/DioSaputra28/belejar-go-dasar/internal/produk/repository"
	productService "github.com/DioSaputra28/belejar-go-dasar/internal/produk/service"
	"github.com/spf13/viper"

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

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"SUPABASE_URL"`
}

func main() {

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("SUPABASE_URL"),
	}

	dbConn := config.DBConn
	if !strings.Contains(dbConn, "sslmode=") {
		if strings.Contains(dbConn, "?") {
			dbConn += "&sslmode=require"
		} else {
			dbConn += "?sslmode=require"
		}
	}

	db, err := database.InitDB(dbConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// Initialize product layer
	productRepo := productRepository.NewProductRepository(db)
	productSvc := productService.NewProductService(productRepo)
	productHandler := handlerProduct.NewProductHandler(productSvc)

	// Initialize category layer
	categoryRepo := categoryRepository.NewCategoryRepository(db)
	categorySvc := categoryService.NewCategoryService(categoryRepo)
	categoryHandler := handlerCategory.NewCategoryHandler(categorySvc)

	// GET localhost:8080/api/produk/{id}
	// PUT localhost:8080/api/produk/{id}
	// DELETE localhost:8080/api/produk/{id}
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			productHandler.GetProductByID(w, r)
		} else if r.Method == "PUT" {
			productHandler.UpdateProduct(w, r)
		} else if r.Method == "DELETE" {
			productHandler.DeleteProduct(w, r)
		}

	})

	// GET localhost:8080/api/produk
	// POST localhost:8080/api/produk
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			productHandler.GetProducts(w, r)
		} else if r.Method == "POST" {
			productHandler.CreateProduct(w, r)
		}

	})

	// GET localhost:8080/api/category/{id}
	// PUT localhost:8080/api/category/{id}
	// DELETE localhost:8080/api/category/{id}
	http.HandleFunc("/api/category/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			categoryHandler.GetCategoryByID(w, r)
		} else if r.Method == "PUT" {
			categoryHandler.UpdateCategory(w, r)
		} else if r.Method == "DELETE" {
			categoryHandler.DeleteCategory(w, r)
		}
	})

	// GET localhost:	
	// POST localhost:8080/api/category
	http.HandleFunc("/api/category", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			categoryHandler.GetCategories(w, r)
		} else if r.Method == "POST" {
			categoryHandler.CreateCategory(w, r)
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
	addr := "0.0.0.0:" + config.Port
	fmt.Println("Server running di", addr)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("gagal running server", err)
	}
}
