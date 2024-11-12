package repository

import (
	"database/sql"
	"errors"
	"math"
	"project/model"
)

type Item struct {
	Db *sql.DB
}

func InitItemRepo(db *sql.DB) *Item {
	return &Item{Db: db}
}

func (repo *Item) Create(item *model.Item) error {
	query := "INSERT INTO items(name, category_id, photo_url, price, purchase_date) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	err := repo.Db.QueryRow(query, item.Name, item.CategoryId, item.PhotoUrl, item.Price, item.PurchaseDate).Scan(&item.Id)

	if err != nil {
		return err
	}
	return nil
}

func (repo *Item) All(criteria model.Search) (int, float64, []model.Item, error) {
	pattern := "%" + criteria.Name + "%"
	var count int
	queryCount := `SELECT COUNT(*)
				FROM items
				WHERE items.name ILIKE $1`
	err := repo.Db.QueryRow(queryCount, pattern).Scan(&count)

	query := `SELECT items.id, items.name , categories.name, photo_url,price,purchase_date,0 AS total_usage_days
				FROM items
				JOIN categories ON categories.id = items.category_id
				WHERE items.deleted_at IS NULL AND items.name ILIKE $1
				LIMIT $2
				OFFSET $3`

	offset := (criteria.Page - 1) * criteria.Limit
	rows, err := repo.Db.Query(query, pattern, criteria.Limit, offset)

	var items []model.Item
	for rows.Next() {
		var item model.Item
		if err := rows.Scan(&item.Id, &item.Name, &item.Category, &item.PhotoUrl, &item.Price, &item.PurchaseDate, &item.TotalUsageDays); err != nil {
			return count, math.Ceil(float64(count) / float64(criteria.Limit)), []model.Item{}, err
		}
		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return count, math.Ceil(float64(count) / float64(criteria.Limit)), []model.Item{}, err
	}
	return count, math.Ceil(float64(count) / float64(criteria.Limit)), items, nil
}

func (repo *Item) Get(id int) (*model.Item, error) {
	query := `SELECT items.id, items.name, categories.name, photo_url,price,purchase_date,current_date-purchase_date
		FROM items JOIN categories ON categories.id = items.category_id
		WHERE items.id = $1 AND items.deleted_at IS NULL`
	var item model.Item
	if err := repo.Db.QueryRow(query, id).Scan(&item.Id, &item.Name, &item.Category, &item.PhotoUrl, &item.Price, &item.PurchaseDate, &item.TotalUsageDays); err != nil {
		return nil, err
	}
	return &item, nil
}

func (repo *Item) Update(item *model.Item) (string, error) {
	//query := "UPDATE items"
	return "", nil
}

func (repo *Item) Delete(id int) error {
	query := `UPDATE items SET deleted_at=NOW() WHERE id = $1 AND deleted_at IS NULL`
	result, err := repo.Db.Exec(query, id)

	if err != nil {
		return err
	}

	nDeleted, err := result.RowsAffected()
	if nDeleted == 0 {
		return errors.New("Barang tidak ditemukan")
	}
	return nil
}
