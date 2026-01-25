package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/DioSaputra28/belejar-go-dasar/internal/category/model"
)


func GetCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.Categories)
}

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