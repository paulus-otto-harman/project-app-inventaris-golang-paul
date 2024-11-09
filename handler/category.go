package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project/model"
	"project/service"
	"project/validation"
	"strings"
)

type CategoryHandler struct {
	CategoryService service.CategoryService
}

func InitCategoryHandler(categoryService service.CategoryService) CategoryHandler {
	return CategoryHandler{CategoryService: categoryService}
}

func (handler CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var category model.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		json.NewEncoder(w).Encode(model.Response{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid input",
		})
		return
	}

	if err := validation.ValidateCategory(category); err != nil {
		json.NewEncoder(w).Encode(model.Response{
			Success:    false,
			StatusCode: http.StatusUnprocessableEntity,
			Message:    err.Error(),
		})
		return
	}

	if err := handler.CategoryService.Create(&category); err != nil {
		errorMessage := fmt.Sprintf("Kategori gagal ditambahkan :: %v", err)
		if strings.Contains(err.Error(), "unique") {
			errorMessage = fmt.Sprintf("Kategori gagal ditambahkan :: Nama Kategori sudah ada")
		}

		json.NewEncoder(w).Encode(model.Response{
			Success:    false,
			StatusCode: http.StatusUnprocessableEntity,
			Message:    errorMessage,
		})
		return
	}

	json.NewEncoder(w).Encode(model.Response{
		Success:    true,
		StatusCode: http.StatusCreated,
		Message:    "Kategori berhasil ditambahkan",
		Data:       category,
	})
}

func (handler CategoryHandler) All(w http.ResponseWriter, r *http.Request) {

}

func (handler CategoryHandler) Get(w http.ResponseWriter, r *http.Request) {

}

func (handler CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {

}

func (handler CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {

}
