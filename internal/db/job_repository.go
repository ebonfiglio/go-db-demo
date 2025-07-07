package db

import (
	"go-db-demo/internal/domain"
	"log"

	"github.com/jmoiron/sqlx"
)

type JobRepository struct {
	db *sqlx.DB
}

func NewJobRepository(db *sqlx.DB) *JobRepository {
	return &JobRepository{db}
}

func (r JobRepository) InsertJob(j *domain.Job, db sqlx.DB) (*domain.Job, error) {
	createdJob := &domain.Job{}
	err := r.db.Get(createdJob,
		`INSERT INTO jobs (name, organization_id)
     VALUES ($1, $2)
     RETURNING id, name, organization_id`,
		j.Name, j.OrganizationID,
	)
	if err != nil {
		log.Fatal(err)
	}
	return createdJob, err
}
