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
	query := `INSERT INTO categories (name,description) VALUES ($1,$2) RETURNING id`

	if err := repo.Db.QueryRow(query, category.Name, &category.Description).Scan(&category.Id); err != nil {
		return err
	}
	return nil
}

func (repo *Category) All() ([]model.Category, error) {
	query := `SELECT categories.id,categories.name, categories.description 
		FROM categories WHERE deleted_at IS NULL`

	rows, err := repo.Db.Query(query)
	var categories []model.Category
	for rows.Next() {
		var category model.Category
		if err := rows.Scan(&category.Id, &category.Name, &category.Description); err != nil {
			return []model.Category{}, err
		}
		categories = append(categories, category)
	}

	if err = rows.Err(); err != nil {
		return []model.Category{}, err
	}
	return categories, nil
}

func (repo *Category) Get(category *model.Category) error {
	query := `SELECT categories.name, categories.description FROM categories WHERE id = $1 AND deleted_at IS NULL`

	if err := repo.Db.QueryRow(query, category.Id).Scan(&category.Name, &category.Description); err != nil {
		return err
	}
	return nil
}

func (repo *Category) Delete(id int) (int64, error) {
	query := `UPDATE categories SET deleted_at=NOW() WHERE id = $1 AND deleted_at IS NULL`
	result, err := repo.Db.Exec(query, id)

	if err != nil {
		return 0, err
	}

	nDeleted, err := result.RowsAffected()
	return nDeleted, nil
}
