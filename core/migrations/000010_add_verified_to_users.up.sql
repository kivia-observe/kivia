ALTER TABLE users ADD COLUMN verified BOOLEAN NOT NULL DEFAULT false;
UPDATE users SET verified = true;
