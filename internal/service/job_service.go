package service

import (
	"fmt"
	"go-db-demo/internal/db"
	"go-db-demo/internal/domain"

	"github.com/jmoiron/sqlx"
)

func CreateJob(j *domain.Job, dbConn *sqlx.DB) (*domain.Job, error) {
	r := db.NewJobRepository(dbConn)

	job, err := r.InsertJob(j)
	if err != nil {
		return nil, fmt.Errorf("could not create job: %w", err)
	}
	return job, err
}

func GetAllJobs(dbConn *sqlx.DB) ([]domain.Job, error) {
	r := db.NewJobRepository(dbConn)

	jobs, err := r.GetAllJobs()
	if err != nil {
		return nil, fmt.Errorf("could not retrieve all jobs: %w", err)
	}

	return jobs, err
}

func GetJob(id int64, dbConn *sqlx.DB) (*domain.Job, error) {
	r := db.NewJobRepository(dbConn)

	job, err := r.GetJob(id)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve job: %w", err)
	}
	return job, err
}

func UpdateJob(j *domain.Job, dbConn *sqlx.DB) (*domain.Job, error) {
	r := db.NewJobRepository(dbConn)

	job, err := r.UpdateJob(j)
	if err != nil {
		return nil, fmt.Errorf("could not update job: %w", err)
	}
	return job, err
}
