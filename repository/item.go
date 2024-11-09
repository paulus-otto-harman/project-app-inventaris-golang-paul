package repository

import (
	"database/sql"
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
