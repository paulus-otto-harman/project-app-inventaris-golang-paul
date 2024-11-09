package repository

import (
	"database/sql"
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
	return nil
}

func (repo *Item) All(criteria model.Search) (int, float64, []model.Item, error) {
	pattern := "%" + criteria.Name + "%"
	//log.Println(pattern)
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
