package db

import (
	"fmt"
	"log"

	"go-db-demo/internal/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect() *sqlx.DB {
	cfg := config.LoadConfig()
	connStr := cfg.Database.GetConnectionString()

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
