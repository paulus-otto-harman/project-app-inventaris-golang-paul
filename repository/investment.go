package repository

import (
	"database/sql"
	"project/model"
)

type Investment struct {
	Db *sql.DB
}

func InitInvestmentRepo(db *sql.DB) *Investment {
	return &Investment{Db: db}
}

func (repo *Investment) All() (*model.Investment, error) {
	query := `WITH RECURSIVE depresiasi(id, price,bulan, rate, d,n) AS (
					SELECT id, price,CAST(EXTRACT(month FROM age(CURRENT_DATE,purchase_date)) AS int), depreciation_rate rate, price-price*depreciation_rate/100 d, 1 n
					FROM items
					UNION ALL
					SELECT id, price,bulan,rate,  d-d*rate/100, n+1 n from depresiasi where n<bulan
				)
				SELECT (SELECT SUM(price) FROM items)::numeric, SUM(terdepresiasi)::numeric
				FROM (SELECT MIN(d) terdepresiasi from depresiasi GROUP BY id)`
	var investment model.Investment
	if err := repo.Db.QueryRow(query).Scan(&investment.TotalInvestment, &investment.DepreciatedValue); err != nil {
		return nil, err
	}
	return &investment, nil
}

func (repo *Investment) Get(id int) (*model.Depreciation, error) {
	query := `WITH RECURSIVE depresiasi(id, price,bulan, rate, d,n) AS (
					SELECT id, price,CAST(EXTRACT(month FROM AGE(CURRENT_DATE,purchase_date)) AS int), depreciation_rate rate, price-price*depreciation_rate/100 d, 1 n
					FROM items
					UNION ALL
					SELECT id, price,bulan,rate,  d-d*rate/100, n+1 n from depresiasi where n<bulan
				)
				SELECT items.id,items.name,items.price::numeric AS initial_price,terdepresiasi::numeric AS depreciated_value,items.depreciation_rate
				FROM (
				SELECT depresiasi.id,MIN(d) terdepresiasi
				FROM depresiasi
				JOIN items ON depresiasi.id=items.id
				WHERE items.id=$1
				GROUP BY depresiasi.id) d
				JOIN items ON d.id=items.id`
	var depreciation model.Depreciation
	if err := repo.Db.QueryRow(query, id).Scan(&depreciation.Id, &depreciation.Name, &depreciation.Price, &depreciation.Value, &depreciation.Rate); err != nil {
		return nil, err
	}
	return &depreciation, nil
}
