package workerservice

import (
	"context"
	"net/http"
	"time"

	"github.com/ysergeyru/revo-testing-task/internal/e"
)

// WorketService is a worker service
type WorkerService struct {
	httpClient *http.Client
}

// New creates a new worker service
func New(httpClient *http.Client) *WorkerService {
	return &WorkerService{
		httpClient: &http.Client{Timeout: time.Second * 3},
	}
}

// BenchURL tests url for max parallel requests it can handle
func (svc *WorkerService) BenchURL(ctx context.Context, url string) e.Error {
	if url == "" {
		return e.NewInternal("no url supplied")
	}

	// TODO implementation

	return nil
}
