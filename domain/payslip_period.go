package domain

import "time"

type PayslipPeriod struct {
	BaseModel

	StartDate    time.Time
	EndDate      time.Time
	PayrollRunAt *time.Time
}
