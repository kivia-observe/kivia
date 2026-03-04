CREATE TABLE logs (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    path VARCHAR(255) NOT NULL, 
    status INTEGER NOT NULL,
    ip_address VARCHAR(20) NOT NULL,
    timestamp TIMESTAMPTZ NOT NULL,
    latency VARCHAR(25) NOT NULL,
    project_Id UUID NOT NULL,
    FOREIGN KEY (project_Id) REFERENCES projects(id)
);