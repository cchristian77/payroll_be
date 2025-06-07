package request

type RunPayroll struct {
	PayrollPeriodID uint64 `json:"payroll_period_id" validate:"required"`
}

type FindPayslipList struct {
	Page            int    `json:"page"`
	PerPage         int    `json:"per_page"`
	PayrollPeriodID uint64 `json:"payroll_period_id"`
	Search          string `json:"search"`
}
