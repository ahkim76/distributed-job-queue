package jobs

import (
	"context"
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

// Get all jobs
func (r *JobRepo) GetAllJobsQuery(ctx context.Context) ([]Job, error) {
	rows, err := r.db.QueryContext(ctx,
	`SELECT
            id,
            queue_name,
            job_type,
            payload,
            status,
            attempts,
            max_attempts,
            priority,
            visible_at,
            lease_expires_at,
            last_error,
            created_at,
            updated_at
        FROM jobs;`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Job 

	for rows.Next() {
		var j Job 

		// map SQL row -> Go variables
		if err := rows.Scan(
            &j.ID,
            &j.QueueName,
            &j.JobType,
            &j.Payload,
            &j.Status,
            &j.Attempts,
            &j.MaxAttempts,
            &j.Priority,
            &j.VisibleAt,
            &j.LeaseExpiresAt,
            &j.LastError,
            &j.CreatedAt,
            &j.UpdatedAt,
        ); err != nil {
            return nil, err
        }
		result = append(result, j)
	}

	return result, nil
}	

// Get job


// Claim job


// Update job


// Get statistics on jobs

