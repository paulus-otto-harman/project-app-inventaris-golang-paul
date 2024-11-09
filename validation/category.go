package validation

import (
	"errors"
	"project/model"
)

func ValidateCategory(category model.Category) error {
	if category.Name == "" {
		return errors.New("Nama kategori tidak boleh kosong")
	}

	return nil
}
