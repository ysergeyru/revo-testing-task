//go:generate mockgen -destination=./mocks.go -source=./repositories.go -package=workerservice

package workerservice

import (
	"context"

	"github.com/ysergeyru/revo-testing-task/internal/model"
)

// JobRepository is a job repository
type JobRepository interface {
	CreateJob(ctx context.Context, job *model.DBJob) error
	UpdateJob(ctx context.Context, job *model.DBJob) error
	GetJobForProcess(ctx context.Context) ([]model.DBJob, error)
}
