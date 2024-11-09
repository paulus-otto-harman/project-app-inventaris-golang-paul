package validation

import (
	"errors"
	"project/model"
)

func ValidateCategory(category model.Category) error {
	if category.Name == "" {
		return errors.New("Nama kategori tidak boleh kosong")
	}

	if len(category.Name) > 20 {
		return errors.New("Nama kategori maksimal 20 karakter")
	}

	if len(category.Description) > 80 {
		return errors.New("Deskripsi kategori maksimal 80 karakter")
	}

	return nil
}
