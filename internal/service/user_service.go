package service

import (
	"fmt"
	"go-db-demo/internal/db"
	"go-db-demo/internal/domain"

	"github.com/jmoiron/sqlx"
)

func CreateUser(u *domain.User, dbConn *sqlx.DB) (*domain.User, error) {
	userRepository := db.NewUserRepository(dbConn)
	user, err := userRepository.InsertUser(u)
	if err != nil {
		return nil, fmt.Errorf("could not create user: %w", err)
	}
	return user, err
}
