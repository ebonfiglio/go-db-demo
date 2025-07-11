package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect() *sqlx.DB {
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=go-db-demo sslmode=disable"

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Could not ping DB:", err)
	}
	fmt.Println("Connected to PostgreSQL!")

	fmt.Println("Ensured 'users' table exists.")

	return db
}
