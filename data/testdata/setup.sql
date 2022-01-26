-- TODO: Run existing migrations instead
-- Or implement some sort of `db:test:prepare`

CREATE TABLE IF NOT EXISTS items(
  id bigserial PRIMARY KEY,
  created_at timestamp(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at timestamp(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
  points int NOT NULL DEFAULT 0,
  title VARCHAR(1024) NOT NULL,
  link VARCHAR(2048),
  from_site VARCHAR(128)
);

CREATE TABLE IF NOT EXISTS users(
  id bigserial PRIMARY KEY,
  created_at timestamp(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at timestamp(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
  karma int NOT NULL DEFAULT(0),
  username varchar(50) NOT NULL,
  email varchar(255),
  hashed_password CHAR(60) NOT NULL
);

INSERT INTO items (title, link, from_site, points) VALUES (
  'Google',
  'https://google.com',
  'google.com',
  10
);

INSERT INTO users (username, email, hashed_password) VALUES (
  'user',
  'user@example.com',
  'hashed-password'
);
