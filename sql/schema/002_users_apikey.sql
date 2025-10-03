-- +goose Up

ALTER TABLE users ADD COLUMN api_key UUID UNIQUE NOT NULL DEFAULT gen_random_uuid();

-- +goose Down
ALTER TABLE users DROP COLUMN api_key;