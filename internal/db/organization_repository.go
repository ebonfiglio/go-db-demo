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
	return createdOrganization, nil
}

func (r OrganizationRepository) GetAll() ([]domain.Organization, error) {
	organizations := make([]domain.Organization, 0)
	err := r.db.Select(&organizations, "select id, name from organizations")
	if err != nil {
		log.Fatal(err)
	}
	return organizations, nil
}

func (r OrganizationRepository) GetOrganization(id int64) (*domain.Organization, error) {
	organization := &domain.Organization{}
	err := r.db.Get(organization, "SELECT id, name FROM organizations WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
	return organization, nil
}

func (r OrganizationRepository) UpdateOrganization(o *domain.Organization) (*domain.Organization, error) {
	updatedOrganization := &domain.Organization{}

	err := r.db.Get(updatedOrganization, "UPDATE organizations SET name = $1 WHERE id = $2 RETURNING id, name", o.Name, o.ID)
	if err != nil {
		log.Fatal(err)
	}

	return updatedOrganization, nil

}

func (r OrganizationRepository) DeleteOrganization(id int64) (int64, error) {
	result, err := r.db.Exec("DELETE organizations WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}

	return result.RowsAffected()
}
