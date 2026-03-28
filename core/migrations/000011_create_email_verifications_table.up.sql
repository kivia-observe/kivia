CREATE TABLE email_verifications (
  id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  email VARCHAR(50) NOT NULL,
  otp_hash VARCHAR(64) NOT NULL,
  attempts INT NOT NULL DEFAULT 0,
  expires_at TIMESTAMPTZ NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  user_id UUID NOT NULL,
  CONSTRAINT fk_email_verifications_users FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX idx_email_verifications_email ON email_verifications(email);
CREATE INDEX idx_email_verifications_user_id ON email_verifications(user_id);
