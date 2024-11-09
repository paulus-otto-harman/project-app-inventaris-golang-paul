package repository

import (
	"database/sql"
)

type Investment struct {
	Db *sql.DB
}

func InitInvestmentRepo(db *sql.DB) *Investment {
	return &Investment{Db: db}
}

func (repo *Investment) All() (interface{}, error) {

	return nil, nil
}

func (repo *Investment) Get(id int) (interface{}, error) {

	return nil, nil
}
