-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE IF NOT EXISTS payslips
(
    id                    BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    created_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by            BIGINT    NOT NULL REFERENCES users (id),
    updated_by            BIGINT REFERENCES users (id),

    user_id               BIGINT    NOT NULL REFERENCES users (id),
    payroll_period_id     BIGINT    NOT NULL REFERENCES payroll_periods (id),
    total_attendance_days INTEGER   NOT NULL,
    total_overtime_days   INTEGER   NOT NULL,
    total_overtime_hours  INTEGER   NOT NULL,
    total_reimbursements  BIGINT    NOT NULL,
    base_salary           BIGINT    NOT NULL,
    total_salary          BIGINT    NOT NULL,

    UNIQUE (user_id, payroll_period_id)
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP TABLE IF EXISTS payslips;