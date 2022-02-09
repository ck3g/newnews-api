CREATE TABLE IF NOT EXISTS auth_sessions(
  id bigserial PRIMARY KEY,
  created_at timestamp(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at timestamp(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
  expired_at timestamp(0) WITH TIME ZONE NOT NULL,
  user_id bigserial NOT NULL,
  token varchar(255) NOT NULL,
  CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX idx_auth_sessions_user_id ON auth_sessions(user_id);
