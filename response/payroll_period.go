package response

import (
	"github.com/cchristian77/payroll_be/domain"
	"time"
)

type PayrollPeriod struct {
	ID        uint64    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	StartDate    string     `json:"start_date"`
	EndDate      string     `json:"end_date"`
	PayrollRunAt *time.Time `json:"payroll_run_at"`
}

func NewPayrollPeriodFromDomain(pp *domain.PayrollPeriod) *PayrollPeriod {
	if pp == nil {
		return nil
	}

	return &PayrollPeriod{
		ID:           pp.ID,
		CreatedAt:    pp.CreatedAt,
		UpdatedAt:    pp.UpdatedAt,
		StartDate:    pp.StartDate.Format(time.DateOnly),
		EndDate:      pp.EndDate.Format(time.DateOnly),
		PayrollRunAt: pp.PayrollRunAt,
	}
}
