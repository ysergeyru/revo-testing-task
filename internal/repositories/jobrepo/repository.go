package objectrepo

import (
	"github.com/ysergeyru/revo-testing-task/internal/pg"
)

// JobRepository is a job repository.
type JobRepository struct {
	db *pg.DB
}

// New creates a new job repository
func New(db *pg.DB) *JobRepository {
	return &JobRepository{
		db: db,
	}
}
