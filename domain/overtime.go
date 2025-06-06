package domain

import "time"

type Overtime struct {
	AttendanceID uint64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	CreatedBy    uint64
	UpdatedBy    *uint64

	Date     time.Time
	Duration uint
	UserID   uint64
}
