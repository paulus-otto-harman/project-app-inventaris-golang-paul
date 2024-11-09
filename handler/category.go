package handler

import (
	"net/http"
	"project/service"
)

type CategoryHandler struct {
	CategoryService service.CategoryService
}

func InitCategoryHandler(categoryService service.CategoryService) CategoryHandler {
	return CategoryHandler{CategoryService: categoryService}
}

func (handler CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {

}
