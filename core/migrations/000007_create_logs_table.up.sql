CREATE TABLE IF NOT EXISTS logs (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    path VARCHAR(255) NOT NULL, 
    status INTEGER NOT NULL,
    ip_address VARCHAR(20) NOT NULL,
    timestamp TIMESTAMPTZ NOT NULL,
    latency VARCHAR(25) NOT NULL,
    api_key_id UUID NOT NULL,
    FOREIGN KEY (api_key_id) REFERENCES api_keys(id)
);