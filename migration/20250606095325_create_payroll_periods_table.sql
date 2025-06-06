-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
-- Payroll Periods
CREATE TABLE IF NOT EXISTS payroll_periods
(
    id             SERIAL PRIMARY KEY,
    created_at     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by     BIGINT NOT NULL REFERENCES users (id),
    updated_by     BIGINT REFERENCES users (id),

    start_date     DATE   NOT NULL,
    end_date       DATE   NOT NULL,
    payroll_run_at TIMESTAMP,

    UNIQUE (start_date, end_date)
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP TABLE IF EXISTS payroll_periods;