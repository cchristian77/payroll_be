-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TYPE reimbursement_status AS ENUM ('PENDING', 'PAID');

CREATE TABLE IF NOT EXISTS reimbursements
(
    id         SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by BIGINT    NOT NULL REFERENCES users (id),
    updated_by BIGINT REFERENCES users (id),

    user_id       BIGINT NOT NULL REFERENCES users (id),
    description   TEXT,
    amount        BIGINT NOT NULL,
    status        reimbursement_status NOT NULL,

    payslip_id    BIGINT REFERENCES payslips (id),
    reimbursed_at TIMESTAMP
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP TABLE IF EXISTS reimbursements;