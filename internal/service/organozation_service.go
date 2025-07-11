package service

import (
	"fmt"
	"go-db-demo/internal/db"
	"go-db-demo/internal/domain"

	"github.com/jmoiron/sqlx"
)

func CreateOrganization(o *domain.Organization, dbConn *sqlx.DB) (*domain.Organization, error) {
	organizationRepository := db.NewOrganizationRepository(dbConn)
	org, err := organizationRepository.InsertOrganization(o)
	if err != nil {
		return nil, fmt.Errorf("could not create organization: %w", err)
	}
	return org, err
}

func GetAllOrganizations(dbConn *sqlx.DB) ([]domain.Organization, error) {
	organizationRepository := db.NewOrganizationRepository(dbConn)
	org, err := organizationRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf("could not retrieve all organizations: %w", err)
	}
	return org, err
}
