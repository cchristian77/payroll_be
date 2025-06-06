package domain

import "time"

type Overtime struct {
	AttendanceID uint64
	CreatedAt    time.Time
	UpdatedAt    time.Time

	Date     time.Time
	Duration uint
	UserID   uint64
}
