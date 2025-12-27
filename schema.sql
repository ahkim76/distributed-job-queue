CREATE TABLE IF NOT EXISTS jobs (
    id                  BIGSERIAL PRIMARY KEY,  -- unique job ID
    queue_name          TEXT NOT NULL DEFAULT 'default', -- lets you support multiple queues (payments, emails, etc)
    job_type            TEXT NOT NULL,          -- which handler/function runs this job
    payload JSONB       NOT NULL,               -- JSON data the job needs

    status              TEXT NOT NULL,          -- job lifecycle: pending | processing | done | dead
    attempts            INT NOT NULL DEFAULT 0, -- how many times job was tried
    max_attempts        INT NOT NULL DEFAULT 5, -- retry cap
    priority            INT NOT NULL DEFAULT 0, 

    visible_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(), -- future-ready: delayed jobs + retry backoff
    lease_expires_at    TIMESTAMPTZ,             -- crash safety: if worker dies, another can reclaim
    last_error          TEXT,                   -- debugging and DLQ info

    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW()
)