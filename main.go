package main

import (
	"log"
	"net/http"
	"project/database"
	"project/router"
)

func main() {
	db := database.DbOpen()
	defer db.Close()
	r := router.NewRouter(db)

	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
