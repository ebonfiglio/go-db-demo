package service

import (
	"fmt"
	"go-db-demo/internal/db"
	"go-db-demo/internal/domain"

	"github.com/jmoiron/sqlx"
)

func CreateOrganization(o *domain.Organization, dbConn *sqlx.DB) (*domain.Organization, error) {
	r := db.NewOrganizationRepository(dbConn)
	org, err := r.InsertOrganization(o)
	if err != nil {
		return nil, fmt.Errorf("could not create organization: %w", err)
	}
	return org, nil
}

func GetAllOrganizations(dbConn *sqlx.DB) ([]domain.Organization, error) {
	r := db.NewOrganizationRepository(dbConn)
	org, err := r.GetAll()
	if err != nil {
		return nil, fmt.Errorf("could not retrieve all organizations: %w", err)
	}
	return org, nil
}

func GetOrganization(id int64, dbConn *sqlx.DB) (*domain.Organization, error) {
	r := db.NewOrganizationRepository(dbConn)
	org, err := r.GetOrganization(id)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve organization: %w", err)
	}
	return org, nil
}

func UpdateOrganization(o *domain.Organization, dbConn *sqlx.DB) (*domain.Organization, error) {
	r := db.NewOrganizationRepository(dbConn)
	org, err := r.UpdateOrganization(o)
	if err != nil {
		return nil, fmt.Errorf("could not update organization: %w", err)
	}
	return org, nil
}
