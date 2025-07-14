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

func (r JobRepository) InsertJob(j *domain.Job) (*domain.Job, error) {
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

func (r JobRepository) GetAllJobs() ([]domain.Job, error) {
	jobs := make([]domain.Job, 0)

	err := r.db.Select(&jobs, "SELECT id, name, organization_id from jobs")
	if err != nil {
		log.Fatal(err)
	}

	return jobs, err
}

func (r JobRepository) GetJob(id int64) (*domain.Job, error) {
	job := &domain.Job{}
	err := r.db.Get(job, "SELECT id, name, organization_id FROM jobs WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}

	return job, err
}
