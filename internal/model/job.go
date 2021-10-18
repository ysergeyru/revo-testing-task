package model

import (
	"time"

	"github.com/volatiletech/null/v8"
)

// DBObject is an object in PostgreSQL.
type DBJob struct {
	tableName         struct{}  `pg:"objects"`
	ID                uint      `pg:"id,notnull,pk"`
	URL               string    `pg:"url,notnull"`
	CanHandleParallel int       `pg:"can_handle_parallel,notnull"`
	CreatedAt         time.Time `pg:"created_at,notnull"`
	ClaimedAt         null.Time `pg:"claimed_at"`
	NextTryAt         null.Time `pg:"next_try_at"`
	TryCount          int       `pg:"try_count"`
	UpdatedAt         time.Time `pg:"updated_at"`
	CompletedAt       null.Time `pg:"completed_at"`
}
