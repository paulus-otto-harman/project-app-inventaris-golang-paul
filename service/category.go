package service

import (
	"project/model"
	"project/repository"
)

type CategoryService struct {
	CategoryRepo repository.Category
}

func InitCategoryService(categoryRepo repository.Category) *CategoryService {
	return &CategoryService{CategoryRepo: categoryRepo}
}

func (categoryService CategoryService) Create(category *model.Category) error {
	return categoryService.CategoryRepo.Create(category)
}

func (categoryService CategoryService) All() ([]model.Category, error) {
	return categoryService.CategoryRepo.All()
}

func (categoryService CategoryService) Get(id int) (*model.Category, error) {
	return categoryService.CategoryRepo.Get(id)
}

func (categoryService CategoryService) Update(category *model.Category) (int, error) {
	return categoryService.CategoryRepo.Update(category)
}

func (categoryService CategoryService) Delete(id int) error {
	return categoryService.CategoryRepo.Delete(id)
}
