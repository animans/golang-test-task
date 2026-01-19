CREATE TABLE IF NOT EXISTS numbers (
  id BIGSERIAL PRIMARY KEY,
  value BIGINT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_numbers_value ON numbers(value);
