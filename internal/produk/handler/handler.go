package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/DioSaputra28/belejar-go-dasar/internal/produk/model"
)

// GetProduk godoc
// @Summary Get all products
// @Description Get list of all products
// @Tags products
// @Produce json
// @Success 200 {array} model.Produk
// @Router /api/produk [get]
func GetProduk(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.Produksi)
}

// CreateProduk godoc
// @Summary Create a new product
// @Description Create a new product with the input payload
// @Tags products
// @Accept json
// @Produce json
// @Param product body model.Produk true "Product to create"
// @Success 201 {object} model.Produk
// @Failure 400 {string} string "Invalid request"
// @Router /api/produk [post]
func CreateProduk(w http.ResponseWriter, r *http.Request) {
	var produkBaru model.Produk
	err := json.NewDecoder(r.Body).Decode(&produkBaru)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	produkBaru.ID = len(model.Produksi) + 1
	model.Produksi = append(model.Produksi, produkBaru)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(produkBaru)
}

// GetProdukByID godoc
// @Summary Get product by ID
// @Description Get a single product by its ID
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} model.Produk
// @Failure 400 {string} string "Invalid Produk ID"
// @Failure 404 {string} string "Produk belum ada"
// @Router /api/produk/{id} [get]
func GetProdukByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	for _, p := range model.Produksi {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

// UpdateProduk godoc
// @Summary Update product
// @Description Update an existing product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body model.Produk true "Product to update"
// @Success 200 {object} model.Produk
// @Failure 400 {string} string "Invalid request"
// @Failure 404 {string} string "Produk belum ada"
// @Router /api/produk/{id} [put]
func UpdateProduk(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	var updateProduk model.Produk
	err = json.NewDecoder(r.Body).Decode(&updateProduk)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	for i := range model.Produksi {
		if model.Produksi[i].ID == id {
			updateProduk.ID = id
			model.Produksi[i] = updateProduk

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateProduk)
			return
		}
	}
	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

// DeleteProduk godoc
// @Summary Delete product
// @Description Delete a product by ID
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "Invalid Produk ID"
// @Failure 404 {string} string "Produk belum ada"
// @Router /api/produk/{id} [delete]
func DeleteProduk(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}
	for i, p := range model.Produksi {
		if p.ID == id {
			model.Produksi = append(model.Produksi[:i], model.Produksi[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "sukses delete",
			})

			return
		}
	}

	http.Error(w, "Produk belum ada", http.StatusNotFound)
}
