package db

import (
	"go-db-demo/internal/domain"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) InsertUser(u *domain.User) (*domain.User, error) {
	createdUser := &domain.User{}
	err := r.db.Get(createdUser,
		"INSERT INTO users (name) VALUES ($1) RETURNING id, name",
		u.Name,
	)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}
