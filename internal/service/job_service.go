package service

import (
	"fmt"
	"go-db-demo/internal/db"
	"go-db-demo/internal/domain"

	"github.com/jmoiron/sqlx"
)

func CreateJob(j *domain.Job, dbConn *sqlx.DB) (*domain.Job, error) {
	r := db.NewJobRepository(dbConn)

	job, err := r.InsertJob(j, *dbConn)
	if err != nil {
		return nil, fmt.Errorf("could not create job: %w", err)
	}
	return job, err
}
