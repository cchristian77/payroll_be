-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE IF NOT EXISTS attendances
(
    id         SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by BIGINT    NOT NULL REFERENCES users (id),
    updated_by BIGINT REFERENCES users (id),

    user_id    BIGINT REFERENCES users (id),
    date       DATE      NOT NULL,
    check_in   TIMESTAMP NOT NULL,
    check_out  TIMESTAMP,
    UNIQUE (user_id, date)
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP TABLE IF EXISTS attendances;