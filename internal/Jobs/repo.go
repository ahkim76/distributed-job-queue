package jobs

import (
	"database/sql"
)

type JobStatus string

type JobRepo struct {
	db *sql.DB
}

func newJobRepo(db *sql.DB) *JobRepo {
	return &JobRepo{db: db}
}

const (
	StatusPending JobStatus = "pending"
)

// DB functions: insert, get, claim, update, stats

// Insert job

// Get job

// Claim job

// Update job

// Get statistics on jobs