package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"project/config"
)

func PgConnect() (*sql.DB, error) {
	return sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=localhost", config.DbUser, config.DbPassword, config.DbName))
}

func DbOpen() *sql.DB {
	db, err := PgConnect()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
