CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

BEGIN;

CREATE TABLE movies (
  id BIGSERIAL PRIMARY KEY,
  title VARCHAR(225),
  imdb_code VARCHAR(50) not NULL,
  views INT,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON movies
FOR EACH ROW
  EXECUTE PROCEDURE trigger_set_timestamp();
COMMIT;
