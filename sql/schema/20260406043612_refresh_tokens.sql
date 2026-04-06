-- +goose Up
CREATE TABLE refresh_tokens (
    token TEXT PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id int,
    expires_at TIMESTAMP,
    revoked_at TIMESTAMP DEFAULT null,
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE
)

-- +goose Down
DROP TABLE refresh_tokens