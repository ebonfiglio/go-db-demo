package service

import (
	"go-db-demo/internal/db"
	"go-db-demo/internal/domain"
	"log"

	"github.com/jmoiron/sqlx"
)

func CreateUser(u *domain.User, dbConn *sqlx.DB) (*domain.User, error) {
	userRepository := db.NewUserRepository(dbConn)
	user, err := userRepository.InsertUser(u)
	if err != nil {
		log.Fatal(err)
	}
	return user, err
}
