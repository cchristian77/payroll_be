-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE IF NOT EXISTS request_logs
(
    id           SERIAL PRIMARY KEY,
    created_at   TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,

    request_id   VARCHAR(255) NOT NULL,
    user_id      BIGINT       NOT NULL REFERENCES users (id),
    activity     VARCHAR(255) NOT NULL,
    entity       VARCHAR(255) NOT NULL,
    reference_id BIGINT       NOT NULL,
    client_ip    VARCHAR(255) NOT NULL
);


-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP TABLE request_logs;