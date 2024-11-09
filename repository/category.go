package repository

import (
	"database/sql"
	"errors"
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

func (repo *Category) Get(id int) (*model.Category, error) {
	query := `SELECT categories.id, categories.name, categories.description
		FROM categories WHERE categories.id = $1 AND deleted_at IS NULL`
	var category model.Category
	if err := repo.Db.QueryRow(query, id).Scan(&category.Id, &category.Name, &category.Description); err != nil {
		return nil, err
	}
	return &category, nil
}

func (repo *Category) Update(category *model.Category) (int, error) {
	query := `UPDATE categories SET name=$1,description=$2 WHERE id = $3 AND deleted_at IS NULL`
	result, err := repo.Db.Exec(query, category.Name, category.Description, category.Id)

	if err != nil {
		return 0, err
	}

	nUpdated, err := result.RowsAffected()
	return int(nUpdated), nil
}

func (repo *Category) Delete(id int) error {
	query := `UPDATE categories SET deleted_at=NOW() WHERE id = $1 AND deleted_at IS NULL`
	result, err := repo.Db.Exec(query, id)

	if err != nil {
		return err
	}

	nDeleted, err := result.RowsAffected()
	if nDeleted == 0 {
		return errors.New("Kategori tidak ditemukan")
	}
	return nil
}
