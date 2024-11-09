package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"project/lib"
	"project/model"
	"project/service"
	"project/validation"
	"strconv"
	"strings"
)

type CategoryHandler struct {
	CategoryService service.CategoryService
}

func InitCategoryHandler(categoryService service.CategoryService) CategoryHandler {
	return CategoryHandler{CategoryService: categoryService}
}

func (handler CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var category model.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		lib.JsonResponse(w).Fail(http.StatusBadRequest, "Invalid input")
		return
	}

	if err := validation.ValidateCategory(category); err != nil {
		lib.JsonResponse(w).Fail(http.StatusUnprocessableEntity, err.Error())
		return
	}

	if err := handler.CategoryService.Create(&category); err != nil {
		errorMessage := err.Error()
		if strings.Contains(errorMessage, "unique") {
			errorMessage = "Nama Kategori sudah ada"
		}
		lib.JsonResponse(w).Fail((http.StatusUnprocessableEntity), fmt.Sprintf("Database Error :: %v", errorMessage))
		return
	}

	lib.JsonResponse(w).Success(http.StatusCreated, "Kategori berhasil ditambahkan", category)
}

func (handler CategoryHandler) All(w http.ResponseWriter, r *http.Request) {
	categories, err := handler.CategoryService.All()
	if err != nil {
		lib.JsonResponse(w).Fail(http.StatusNotFound, err.Error())
		return
	}
	lib.JsonResponse(w).Success(http.StatusOK, "", categories)
}

func (handler CategoryHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		lib.JsonResponse(w).Fail(http.StatusBadRequest, "Invalid Category ID")
		return
	}

	category, err := handler.CategoryService.Get(id)
	if err != nil {
		lib.JsonResponse(w).Fail(http.StatusNotFound, "Kategori tidak ditemukan")
		return
	}
	lib.JsonResponse(w).Success(http.StatusOK, "", category)
}

func (handler CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		lib.JsonResponse(w).Fail(http.StatusBadRequest, "Invalid Category ID")
		return
	}

	category := model.Category{Id: id}
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		lib.JsonResponse(w).Fail(http.StatusBadRequest, "Invalid input")
		return
	}

	if err := validation.ValidateCategory(category); err != nil {
		lib.JsonResponse(w).Fail(http.StatusUnprocessableEntity, err.Error())
		return
	}

	updated, err := handler.CategoryService.Update(&category)
	if err != nil {
		errorMessage := err.Error()
		if strings.Contains(errorMessage, "unique") {
			errorMessage = "Nama Kategori sudah ada"
		}
		lib.JsonResponse(w).Fail((http.StatusUnprocessableEntity), fmt.Sprintf("Database Error :: %v", errorMessage))
		return
	}

	if updated == 0 {
		lib.JsonResponse(w).Fail(http.StatusNotFound, "Kategori tidak ditemukan")
	}

	lib.JsonResponse(w).Success(http.StatusCreated, "Kategori berhasil diperbarui", category)
}

func (handler CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		lib.JsonResponse(w).Fail(http.StatusBadRequest, "Invalid Category ID")
		return
	}

	if err := handler.CategoryService.Delete(id); err != nil {
		if strings.Contains(err.Error(), "tidak ditemukan") {
			lib.JsonResponse(w).Fail(http.StatusNotFound, "Kategori tidak ditemukan")
		}
		lib.JsonResponse(w).Fail(http.StatusBadRequest, err.Error())
		return
	}
	lib.JsonResponse(w).Success(http.StatusNoContent, "Kategori tidak ditemukan", nil)
}
