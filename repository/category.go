package repository

import (
	"database/sql"
	"project/model"
)

type Category struct {
	Db *sql.DB
}

func InitCategoryRepo(db *sql.DB) *Category {
	return &Category{Db: db}
}

func (repo *Category) Create(category *model.Category) error {
	return nil
}
