package db

import (
	"go-db-demo/internal/domain"
	"log"

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
		log.Fatal(err)
	}
	return createdUser, nil
}

func (r UserRepository) GetAllUsers() ([]domain.User, error) {
	users := make([]domain.User, 0)
	err := r.db.Select(&users, "SELECT id, name, job_id, organization_id from users")
	if err != nil {
		log.Fatal(err)
	}

	return users, nil
}

func (r UserRepository) GetUser(id int64) (*domain.User, error) {
	user := &domain.User{}
	err := r.db.Get(user, "SELECT id, name, job_id, organization_id FROM users WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}

	return user, nil
}

func (r UserRepository) UpdateUser(u *domain.User) (*domain.User, error) {
	updatedUser := &domain.User{}
	err := r.db.Get(updatedUser, "UPDATE users set name = $1, job_id = $2, organization_id = $3 WHERE id = $4 RETURNING id, name, job_id, organization_id", u.Name, u.JobID, u.OrganizationID, u.ID)
	if err != nil {
		log.Fatal(err)
	}
	return updatedUser, nil
}
