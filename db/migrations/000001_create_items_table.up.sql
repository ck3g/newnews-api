CREATE TABLE IF NOT EXISTS items(
  id bigserial PRIMARY KEY,
  created_at timestamp(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at timestamp(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
  points int NOT NULL DEFAULT 0,
  title VARCHAR(1024) NOT NULL,
  link VARCHAR(2048),
  from_site VARCHAR(128)
);
