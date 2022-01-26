CREATE TABLE IF NOT EXISTS users(
  id bigserial PRIMARY KEY,
  created_at timestamp(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at timestamp(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
  karma int NOT NULL DEFAULT(0),
  username varchar(50) NOT NULL,
  email varchar(255),
  hashed_password CHAR(60) NOT NULL
);

CREATE UNIQUE INDEX idx_users_username ON users(username);
CREATE UNIQUE INDEX idx_users_email ON users(email);
