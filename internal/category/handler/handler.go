package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/DioSaputra28/belejar-go-dasar/internal/category/model"
)

// GetCategory godoc
// @Summary Get all categories
// @Description Get list of all categories
// @Tags categories
// @Produce json
// @Success 200 {array} model.Category
// @Router /api/category [get]
func GetCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.Categories)
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new category with the input payload
// @Tags categories
// @Accept json
// @Produce json
// @Param category body model.Category true "Category to create"
// @Success 201 {object} model.Category
// @Failure 400 {string} string "Invalid request"
// @Router /api/category [post]
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category model.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	category.ID = len(model.Categories) + 1
	model.Categories = append(model.Categories, category)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}

// GetCategoryByID godoc
// @Summary Get category by ID
// @Description Get a single category by its ID
// @Tags categories
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} model.Category
// @Failure 400 {string} string "Invalid Category ID"
// @Failure 404 {string} string "Category not found"
// @Router /api/category/{id} [get]
func GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	for i := range model.Categories {
		if model.Categories[i].ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(model.Categories[i])
			return
		}
	}

	http.Error(w, "Category not found", http.StatusNotFound)
}

// UpdateCategory godoc
// @Summary Update category
// @Description Update an existing category by ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param category body model.Category true "Category to update"
// @Success 200 {object} model.Category
// @Failure 400 {string} string "Invalid request"
// @Failure 404 {string} string "Category not found"
// @Router /api/category/{id} [put]
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	var updateCategory model.Category
	err = json.NewDecoder(r.Body).Decode(&updateCategory)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	for i := range model.Categories {
		if model.Categories[i].ID == id {
			updateCategory.ID = id
			model.Categories[i] = updateCategory

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateCategory)
			return
		}
	}

	http.Error(w, "Category not found", http.StatusNotFound)
}

// DeleteCategory godoc
// @Summary Delete category
// @Description Delete a category by ID
// @Tags categories
// @Produce json
// @Param id path int true "Category ID"
// @Success 204 "No Content"
// @Failure 400 {string} string "Invalid Category ID"
// @Failure 404 {string} string "Category not found"
// @Router /api/category/{id} [delete]
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	for i := range model.Categories {
		if model.Categories[i].ID == id {
			model.Categories = append(model.Categories[:i], model.Categories[i+1:]...)

			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Category not found", http.StatusNotFound)
}
