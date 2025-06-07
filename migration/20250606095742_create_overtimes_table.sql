-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS overtimes
(
    attendance_id BIGINT PRIMARY KEY REFERENCES attendances (id),
    created_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by    BIGINT    NOT NULL REFERENCES users (id),
    updated_by    BIGINT REFERENCES users (id),

    user_id       BIGINT REFERENCES users (id),
    date          DATE      NOT NULL,
    duration      SMALLINT  NOT NULL
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP TABLE IF EXISTS overtimes;
