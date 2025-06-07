package response

import "time"

type PayrollPeriod struct {
	ID        uint64    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	StartDate    string     `json:"start_date"`
	EndDate      string     `json:"end_date"`
	PayrollRunAt *time.Time `json:"payroll_run_at"`
}
