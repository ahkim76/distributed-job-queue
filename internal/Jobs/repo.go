package jobs

import (
	"database/sql"
)

type JobStatus string

type JobRepo struct {
	db *sql.DB
}

func NewJobRepo(db *sql.DB) *JobRepo {
	return &JobRepo{db: db}
}

const (
	StatusPending JobStatus = "pending"
	StatusComplete JobStatus = "complete"
	StatusFailed JobStatus = "failed"
	StatusProcessing JobStatus = "processing"
	StatusNew JobStatus = "new"
)

// DB functions: insert, get, claim, update, stats

// Insert job

// Get job

// Claim job

// Update job

// Get statistics on jobs