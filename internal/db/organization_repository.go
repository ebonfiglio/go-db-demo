package db

import (
	"go-db-demo/internal/domain"
	"log"

	"github.com/jmoiron/sqlx"
)

type OrganizationRepository struct {
	db *sqlx.DB
}

func NewOrganizationRepository(db *sqlx.DB) *OrganizationRepository {
	return &OrganizationRepository{db}
}

func (r OrganizationRepository) InsertOrganization(o *domain.Organization) (*domain.Organization, error) {
	createdOrganization := &domain.Organization{}
	err := r.db.Get(createdOrganization,
		"INSERT INTO organizations (name) VALUES ($1) RETURNING id, name",
		o.Name,
	)
	if err != nil {
		log.Fatal(err)
	}
	return createdOrganization, err
}

func (r OrganizationRepository) GetAll() ([]domain.Organization, error) {
	organizations := make([]domain.Organization, 0)
	err := r.db.Select(&organizations, "select id, name from organizations")
	if err != nil {
		log.Fatal(err)
	}
	return organizations, err
}
