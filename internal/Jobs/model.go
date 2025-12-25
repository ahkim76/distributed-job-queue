package jobs

import (
	"encoding/json"
	"time"
)

// What a Job struct looks like in Go. The model should represent one job row in Postgres

type Job struct {
	ID int64 `json:"id"`
	QueueName string `json:"queue_name"`
	JobType string `json:"job_type"`
	Payload json.RawMessage `json:"payload"`
	Status string `json:"status"`
	Attempts int `json:"attempts"`
	MaxAttempts int `json:"max_attempts"`
	Priority int `json:"priority"`
	VisibleAt time.Time `json:"visible_at"`
	LeaseExpiresAt *time.Time `json:"lease_expires_at,omitempty"`
	LastError *string `json:"last_error,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}