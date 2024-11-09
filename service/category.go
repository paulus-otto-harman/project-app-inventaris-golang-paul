package service

import (
	"project/model"
	"project/repository"
)

type CategoryService struct {
	CategoryRepo repository.Category
}

func InitCategoryService(repo repository.Category) *CategoryService {
	return &CategoryService{CategoryRepo: repo}
}

func (categoryService CategoryService) Create(category *model.Category) error {
	return categoryService.CategoryRepo.Create(category)
}
