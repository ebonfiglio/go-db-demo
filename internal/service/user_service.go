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

func GetAllUsers(dbConn *sqlx.DB) ([]domain.User, error) {
	r := db.NewUserRepository(dbConn)

	users, err := r.GetAllUsers()

	if err != nil {
		return nil, fmt.Errorf("could not retrieve all users: %w", err)
	}

	return users, err
}

func GetUser(id int64, dbConn sqlx.DB) (*domain.User, error) {
	r := db.NewUserRepository(&dbConn)

	user, err := r.GetUser(id)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve user: %w", err)
	}
	return user, err
}

func UpdateUser(u *domain.User, dbConn *sqlx.DB) (*domain.User, error) {
	r := db.NewUserRepository(dbConn)

	user, err := r.UpdateUser(u)
	if err != nil {
		return nil, fmt.Errorf("could not update user: %w", err)
	}
	return user, err
}
