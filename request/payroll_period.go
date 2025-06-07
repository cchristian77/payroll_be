package request

type UpsertPayrollPeriod struct {
	ID uint64 `json:"id"`

	StartDate string `json:"start_date" validate:"required"`
	EndDate   string `json:"end_date" validate:"required"`
}
