package repository

import "database/sql"

type Investment struct {
	Db *sql.DB
}

func InitInvestmentRepo(db *sql.DB) *Investment {
	return &Investment{Db: db}
}
