CREATE TABLE refresh_tokens (
  id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  token varchar UNIQUE NOT NULL,
  expires_at timestamptz NOT NULL,
  blacklisted boolean NOT NULL,
  user_id uuid NOT NULL,
  created_at timestamptz DEFAULT now()
);

ALTER TABLE refresh_tokens
  ADD CONSTRAINT fk_refresh_tokens_users
  FOREIGN KEY (user_id)
  REFERENCES users(id)
  ON DELETE CASCADE;