package domain

import "time"

type PayrollPeriod struct {
	BaseModel

	StartDate    time.Time
	EndDate      time.Time
	PayrollRunAt *time.Time
}
