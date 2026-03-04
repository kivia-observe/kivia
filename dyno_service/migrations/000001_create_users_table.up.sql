CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
  id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  name varchar(20) NOT NULL,
  email varchar(50) UNIQUE NOT NULL,
  profile_picture varchar(100),
  joined_at timestamptz NOT NULL DEFAULT now(),
  password varchar(100) NOT NULL
);
