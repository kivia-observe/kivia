CREATE TABLE api_keys (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    key TEXT NOT NULL,
    user_id UUID NOT NULL,
    project_id UUID NOT NULL,
    revoked BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE
);