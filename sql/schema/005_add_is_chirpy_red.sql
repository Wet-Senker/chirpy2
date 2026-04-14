-- +goose Up
ALTER TABLE users
ADD COLUMN is_chirpy BOOLEAN NOT NULL DEFAULT FALSE;

-- +goose Down
ALTER TABLE users
DROP COLUMN is_chirpy;